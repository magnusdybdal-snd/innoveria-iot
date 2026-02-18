/**
 *
 */
export type Gateway = {
  id: string;
  name: string;
  online?: boolean;
  euid: string;
  lastSeen: string;
};

export const mockGateways: Gateway[] = [
  {
    id: "1",
    name: "Gateway 1",
    online: true,
    euid: "A1B2C3D4E5F6",
    lastSeen: "2 min ago",
  },
  {
    id: "2",
    name: "Gateway 2",
    online: false,
    euid: "B2C3D4E5F6A1",
    lastSeen: "1 hour ago",
  },
  {
    // Gateway without status (should return false status)
    id: "3",
    name: "Gateway 3",
    euid: "B2D3D5E5A6A1",
    lastSeen: "4 days ago",
  },
];
