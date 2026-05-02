/**
 * HTTP Load Test — innoveria-iot
 *
 * Maps to requirements:
 *   Req 1: 300 concurrent users
 *   Req 3: Response time < 1s (p95) under normal conditions
 *
 * Usage:
 *   k6 run --env K6_USERNAME=user@example.com --env K6_PASSWORD=secret \
 *          --env BASE_URL=http://localhost:8081 \
 *          docs/tests/k6/http-load-test.js
 *
 * Scenarios:
 *   baseline       — 10 VUs for 3 min. Validates Req 3 at low load.
 *   concurrent_300 — ramp to 300 VUs, hold 5 min, ramp down. Validates Req 1 + 3.
 *
 * To run only one scenario:
 *   k6 run --scenario baseline ...
 *   k6 run --scenario concurrent_300 ...
 */

import http from 'k6/http';
import { sleep, check, group } from 'k6';
import { Rate, Trend } from 'k6/metrics';

const errorRate = new Rate('errors');
const sensorLatestDuration = new Trend('sensor_latest_duration', true);

const BASE_URL = __ENV.BASE_URL || 'http://localhost:8081';
const USERNAME = __ENV.K6_USERNAME;
const PASSWORD = __ENV.K6_PASSWORD;

export const options = {
  scenarios: {
    concurrent_300: {
      executor: 'ramping-vus',
      startVUs: 0,
      stages: [
        { duration: '2m', target: 300 },  // ramp to requirement
        { duration: '2m', target: 600 },  // 2× headroom
        { duration: '2m', target: 1000 }, // stress ceiling
        { duration: '3m', target: 1000 }, // hold
        { duration: '1m', target: 0 },    // ramp down
      ],
      tags: { scenario: 'concurrent_300' },
    },
  },
  thresholds: {
    'http_req_duration': ['p(95)<1000'],
    'http_req_failed': ['rate<0.01'],
    'errors': ['rate<0.01'],
  },
};

export function setup() {
  if (!USERNAME || !PASSWORD) {
    throw new Error('K6_USERNAME and K6_PASSWORD environment variables are required');
  }

  const loginRes = http.post(
    `${BASE_URL}/api/v1/auth/login`,
    JSON.stringify({ email: USERNAME, password: PASSWORD }),
    { headers: { 'Content-Type': 'application/json' } },
  );

  if (!check(loginRes, { 'setup: login 200': (r) => r.status === 200 })) {
    throw new Error(`Login failed (${loginRes.status}): ${loginRes.body}`);
  }

  const token = loginRes.json('access_token');

  // Fetch sensor EUIs once so VUs can request their latest readings
  const sensorsRes = http.get(`${BASE_URL}/api/v1/device/sensors`, {
    headers: { Authorization: `Bearer ${token}` },
  });

  const sensorEUIs = sensorsRes.status === 200
    ? (sensorsRes.json('sensors') || []).map((s) => s.device_eui).filter(Boolean)
    : [];

  console.log(`Setup complete. Token acquired. ${sensorEUIs.length} sensor(s) found.`);

  return { token, sensorEUIs };
}

export default function ({ token, sensorEUIs }) {
  const headers = { Authorization: `Bearer ${token}` };

  // Simulates a user loading their profile
  group('user_info', () => {
    const res = http.get(`${BASE_URL}/api/v1/auth/me`, { headers });
    check(res, { 'me: 200': (r) => r.status === 200 });
    errorRate.add(res.status !== 200);
  });

  // Simulates loading the sensor dashboard
  group('sensor_dashboard', () => {
    const sensorsRes = http.get(`${BASE_URL}/api/v1/device/sensors`, { headers });
    check(sensorsRes, { 'sensors: 200': (r) => r.status === 200 });
    errorRate.add(sensorsRes.status !== 200);

    // Request the latest reading for a random sensor
    if (sensorEUIs.length > 0) {
      const eui = sensorEUIs[Math.floor(Math.random() * sensorEUIs.length)];
      const start = Date.now();
      const latestRes = http.get(
        `${BASE_URL}/api/v1/collection/latest?device_eui=${eui}`,
        { headers },
      );
      sensorLatestDuration.add(Date.now() - start);
      check(latestRes, { 'latest reading: 200': (r) => r.status === 200 });
      errorRate.add(latestRes.status !== 200);
    }
  });

  // Simulates loading the production overview
  group('production_view', () => {
    const prRes = http.get(`${BASE_URL}/api/v1/erp/production-resources`, { headers });
    check(prRes, { 'production-resources: 200': (r) => r.status === 200 });
    errorRate.add(prRes.status !== 200);

    const ordersRes = http.get(`${BASE_URL}/api/v1/context/orders`, { headers });
    check(ordersRes, { 'orders: 200': (r) => r.status === 200 });
    errorRate.add(ordersRes.status !== 200);
  });

  // Think time between iterations — models a user reading the page
  sleep(Math.random() * 3 + 1);
}
