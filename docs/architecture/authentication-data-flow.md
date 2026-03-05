flowchart LR

Client[Frontend]

subgraph Gateway["API Gateway"]

JWT["JWT Validation"]

CompanyCtx["Resolve Company Context
(active_company_id)"]

RBACCheck["RBAC Authorization
Evaluate:
(user_id, company_id, permission)"]

InjectHeaders["Inject Trusted Context
X-User-Id
X-Company-Id
X-Roles
X-Permissions
X-Request-Id"]

end


subgraph AuthService["Auth Service"]

subgraph AuthHandlers["Handler Layer"]
Login["POST /auth/login"]
Me["GET /auth/me"]
SwitchCompany["POST /auth/switch-company"]
CreateCompany["POST /companies (Platform Admin)"]
end

subgraph AuthServices["Service Layer"]

AuthLogic["Auth Service"]

Onboarding["Company Onboarding"]

end


subgraph AuthRepos["Repository Layer"]

UserRepo["users"]

CompanyRepo["companies"]

MembershipRepo["user_company_memberships"]

RoleRepo["roles"]

PermissionRepo["permissions"]

end

AuthDB[(Auth Database)]

end



subgraph DeviceService["Device Service"]

DeviceHandler["Handler"]

DeviceServiceLogic["Service"]

DeviceRepo["Repository"]

DeviceDB[(Device Database)]

end


subgraph CollectionService["Collection Service"]

MeasurementHandler["Handler"]

MeasurementService["Service"]

MeasurementRepo["Repository"]

CollectionDB[(Collection Database)]

end



subgraph Provisioning["Onboarding Service"]

CompanyCreated["CompanyCreated Event"]

end



subgraph External["ChirpStack Service"]

ChirpStack["ChirpStack API"]

end



Client --> JWT
JWT --> CompanyCtx
CompanyCtx --> RBACCheck
RBACCheck --> InjectHeaders


InjectHeaders --> Login
InjectHeaders --> Me
InjectHeaders --> SwitchCompany
InjectHeaders --> DeviceHandler
InjectHeaders --> MeasurementHandler


Login --> AuthLogic
Me --> AuthLogic
SwitchCompany --> AuthLogic
CreateCompany --> Onboarding


AuthLogic --> UserRepo
AuthLogic --> MembershipRepo
AuthLogic --> RoleRepo
AuthLogic --> PermissionRepo


Onboarding --> CompanyRepo
Onboarding --> MembershipRepo
Onboarding --> RoleRepo
Onboarding --> CompanyCreated


UserRepo --> AuthDB
CompanyRepo --> AuthDB
MembershipRepo --> AuthDB
RoleRepo --> AuthDB
PermissionRepo --> AuthDB


CompanyCreated --> DeviceServiceLogic


DeviceHandler --> DeviceServiceLogic
DeviceServiceLogic --> DeviceRepo
DeviceRepo --> DeviceDB
DeviceServiceLogic --> ChirpStack


MeasurementHandler --> MeasurementService
MeasurementService --> MeasurementRepo
MeasurementRepo --> CollectionDB
