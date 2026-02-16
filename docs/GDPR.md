# Fabrikkpuls — GDPR Notes

## 1. Is GDPR Relevant for Fabrikkpuls?

Yes. Although Fabrikkpuls primarily handles industrial sensor data (temperature, nitrogen consumption, electric current), the system also processes personal data through user accounts. GDPR applies to all processing of personal data within the EU/EEA, regardless of whether it is the primary purpose of the system. ([GDPR Article 2 — Material scope](https://gdpr-info.eu/art-2-gdpr/))

IoT systems like Fabrikkpuls have specific privacy considerations because they collect data continuously, and combinations of data streams can potentially identify individuals indirectly. ([Legal IT Group — GDPR and Internet of Things](https://legalitgroup.com/en/gdpr-and-internet-of-things-iot/))

---

## 2. What Data Constitutes Personal Data in Fabrikkpuls?

### Clearly personal data (in auth-db)

| Data | Why it is personal data |
|------|------------------------|
| Name | Directly identifiable |
| Email | Directly identifiable |
| PasswordHash | Linked to identifiable person |
| Role | Linked to identifiable person |
| LastLoggedIn (timestamp) | Behavioural data linked to person |
| IP addresses (if logged) | Considered personal data under GDPR |

### Not personal data (in collection-db, erp-db, context-db)

Sensor data (temperature, nitrogen consumption, electric current), order data (order number, product name, quantity), machine information, and context data are **not personal data** in themselves. No individual can be identified from a temperature reading on machine X.

### Grey area: indirect identification

If sensor data from an individual workstation can be combined with shift schedules from the ERP system (e.g., "machine X at 14:00" + "employee Y was operating machine X at 14:00"), the sensor data could **indirectly** become personal data. GDPR defines personal data as any information that can be linked to an identifiable person, including indirect identification. ([GDPR Article 4(1) — Definitions](https://gdpr-info.eu/art-4-gdpr/))

For Fabrikkpuls's MVP this is unlikely, but it should be noted as a consideration for production deployment.

---

## 3. Relevant GDPR Articles for Fabrikkpuls

### Article 5 — Principles relating to processing of personal data

All personal data must be processed in accordance with the following principles ([GDPR Article 5](https://gdpr-info.eu/art-5-gdpr/)):

- **Lawfulness, fairness and transparency** — users must know what is collected and why
- **Purpose limitation** — data is collected only for specific, explicit purposes
- **Data minimisation** — only necessary personal data is collected
- **Accuracy** — data must be correct and up to date
- **Storage limitation** — data must not be stored longer than necessary
- **Integrity and confidentiality** — appropriate security measures must protect the data
- **Accountability** — the controller must be able to demonstrate compliance

**For Fabrikkpuls:** We only collect what is necessary for user accounts (name, email, role). We do not store unnecessary personal information. Sensor data does not contain personal data.

### Article 6 — Lawfulness of processing

Processing of personal data is only lawful if at least one of the conditions in Article 6(1) is met. ([GDPR Article 6](https://gdpr-info.eu/art-6-gdpr/))

**For Fabrikkpuls:** The most relevant legal bases are:
- **Article 6(1)(b) — Performance of a contract**: The user needs an account to use the system. Processing of name, email, and role is necessary to deliver the service.
- **Article 6(1)(f) — Legitimate interest**: Logging of login timestamps and IP addresses for security and troubleshooting purposes.

### Article 17 — Right to erasure ("right to be forgotten")

Data subjects have the right to have their personal data erased without undue delay. ([GDPR Article 17](https://gdpr-info.eu/art-17-gdpr/))

**For Fabrikkpuls:** We must be able to delete a user account upon request. This has cascading effects that must be handled:
- Dashboard configuration linked to the user → deleted
- Action log / audit trail → anonymised (replace userId with "deleted user")
- User removed from auth-db

Sensor data and context data are not linked to individual users and are unaffected.

### Article 25 — Data protection by design and by default

Privacy and data protection must be integrated into the system from the design stage, not added afterwards. The controller must implement appropriate technical and organisational measures to ensure that only necessary personal data is processed by default. ([GDPR Article 25](https://gdpr-info.eu/art-25-gdpr/), [EDPB Guidelines 4/2019 on Article 25](https://www.edpb.europa.eu/sites/default/files/files/file1/edpb_guidelines_201904_dataprotection_by_design_and_by_default_v2.0_en.pdf))

**How Fabrikkpuls implements this:**
- **Data minimisation** — auth-db stores only what is needed (name, email, role, passwordHash). No superfluous personal information.
- **Separation** — personal data is isolated in auth-db. Sensor data in collection-db contains no personal data. The microservice architecture ensures that services that do not need personal data have no access to it.
- **Access control** — role-based access (Factory Worker, Superuser, Admin) restricts who can see what.
- **Password storage** — passwords hashed with encryption, never stored in plaintext.
- **Multi-tenancy** — `selskapId` on all tables ensures data from one company is inaccessible to another.

### Article 32 — Security of processing

The controller must implement appropriate technical and organisational measures to ensure a level of security appropriate to the risk, including encryption and pseudonymisation. ([GDPR Article 32](https://gdpr-info.eu/art-32-gdpr/))

**For Fabrikkpuls:**
- HTTPS between frontend and API Gateway (encrypted transport)
- JWT tokens with short expiry for authentication
- Passwords hashed with encryption
- Database-per-service architecture limits the blast radius of a potential breach — compromising erp-db does not grant access to auth-db
- Separate database users per service — no service has access to another service's data
- Multi-tenancy isolation with `selskapId` prevents cross-contamination between companies

### Article 33 — Notification of a personal data breach to the supervisory authority

In the case of a personal data breach, the supervisory authority must be notified without undue delay and no later than 72 hours after the controller becomes aware of it. All breaches must be documented. ([GDPR Article 33](https://gdpr-info.eu/art-33-gdpr/))

**For Fabrikkpuls:** In a production deployment, there must be a procedure for handling personal data breaches. For the bachelor project, it is sufficient to document that this requirement exists and describe a high-level procedure.

### Article 35 — Data Protection Impact Assessment (DPIA)

When processing, in particular using new technologies, is likely to result in a high risk to the rights and freedoms of natural persons, the controller must carry out a DPIA before the processing begins. ([GDPR Article 35](https://gdpr-info.eu/art-35-gdpr/))

The Article 29 Working Party (now EDPB) has specifically stated that IoT applications could have a significant impact on individuals' daily lives and privacy, and therefore require a DPIA. ([WP29 Guidelines on DPIA, WP248 rev.01](https://gdprhub.eu/Article_35_GDPR))

**For Fabrikkpuls:** Although sensor data itself is not personal data, the system uses new IoT technology and processes user data. A simplified DPIA should be considered. For the bachelor project, it is sufficient to:
- Identify which personal data is processed (user accounts in auth-db)
- Assess the risks (data loss, unauthorised access)
- Document the security measures implemented
- Conclude that the risk is low due to the limited scope of personal data

---

## 4. Practical Measures for Fabrikkpuls

### Must do (for the project)

| Measure | GDPR basis | Status |
|---------|-----------|--------|
| Hash passwords with encryption | Art. 32 (security) | |
| HTTPS between frontend and API Gateway | Art. 32 (encryption in transit) | |
| Role-based access control | Art. 25 (privacy by default) | |
| Multi-tenancy with selskapId on all tables | Art. 25 (access restriction) | |
| Ability to delete user accounts | Art. 17 (right to erasure) | |
| Short expiry on JWT tokens | Art. 32 (security) | |
| Do not log personal data unnecessarily | Art. 5(1)(c) (data minimisation) | |

### Should mention in the report (but not necessarily fully implemented)

| Measure | GDPR basis | Comment |
|---------|-----------|---------|
| Privacy policy | Art. 13/14 (information obligation) | Should be written, can be simplified |
| DPIA | Art. 35 (impact assessment) | Describe as required for production |
| Breach notification procedure | Art. 33 (72-hour deadline) | Describe high-level procedure |
| Data processing agreement with ERP provider | Art. 28 (processor) | Relevant for production |
| Records of processing activities | Art. 30 (processing overview) | Describe as a requirement |

---

## 5. What Makes Fabrikkpuls Low-Risk from a GDPR Perspective

- **The primary purpose is industrial monitoring**, not personal monitoring. Sensor data (temperature, nitrogen, electric current) is not personal data.
- **Limited scope of personal data** — only user accounts (name, email, role). No sensitive personal data (health, biometrics, political beliefs, etc.).
- **B2B context** — users are factory employees, not consumers. No minors.
- **The multi-tenancy architecture** provides built-in data isolation between companies.
- **The microservice architecture** ensures personal data is isolated in auth-db and not spread across the entire system.

---

## 6. Potential Risks to Be Aware Of

- **Indirect identification via sensor data + shift schedules**: If sensor data from individual machines can be linked with who operated the machine (from ERP), this could constitute indirect personal data. Mitigation: avoid importing individual shift schedules/employee data from ERP.
- **Logging and audit trails**: If the system logs user actions (who configured a sensor, who changed a threshold), this is personal data. Mitigation: implement log rotation and anonymisation of old logs.
- **JWT tokens containing personal data**: If tokens contain userId, email, or role, personal data is transported throughout the entire system. Mitigation: keep tokens minimal and with short expiry, and do not log them.

---

## References

### Official GDPR Articles
- [Article 4 — Definitions](https://gdpr-info.eu/art-4-gdpr/)
- [Article 5 — Principles relating to processing](https://gdpr-info.eu/art-5-gdpr/)
- [Article 6 — Lawfulness of processing](https://gdpr-info.eu/art-6-gdpr/)
- [Article 17 — Right to erasure](https://gdpr-info.eu/art-17-gdpr/)
- [Article 25 — Data protection by design and by default](https://gdpr-info.eu/art-25-gdpr/)
- [Article 32 — Security of processing](https://gdpr-info.eu/art-32-gdpr/)
- [Article 33 — Notification of a personal data breach](https://gdpr-info.eu/art-33-gdpr/)
- [Article 35 — Data Protection Impact Assessment (DPIA)](https://gdpr-info.eu/art-35-gdpr/)

### Guidelines and Recommendations
- [EDPB Guidelines 4/2019 on Article 25 — Data Protection by Design and by Default](https://www.edpb.europa.eu/sites/default/files/files/file1/edpb_guidelines_201904_dataprotection_by_design_and_by_default_v2.0_en.pdf)
- [WP29 Guidelines on DPIA (WP248 rev.01)](https://gdprhub.eu/Article_35_GDPR)
- [ICO — When do we need to do a DPIA?](https://ico.org.uk/for-organisations/uk-gdpr-guidance-and-resources/accountability-and-governance/data-protection-impact-assessments-dpias/when-do-we-need-to-do-a-dpia/)

### IoT-Specific Sources
- [Legal IT Group — GDPR and Internet of Things (IoT)](https://legalitgroup.com/en/gdpr-and-internet-of-things-iot/)
- [AWS — Database-per-service pattern](https://docs.aws.amazon.com/prescriptive-guidance/latest/modernization-data-persistence/database-per-service.html) (relevant for architectural isolation of personal data)
