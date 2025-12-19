---
title: Cloud Native Security Controls Catalog
sidebar_position: 1
---

# Cloud Native Security Controls Catalog

The Cloud Native Security Controls Catalog provides comprehensive guidance for securing cloud-native applications and workloads.

## Table of Contents

- [Access](#access)
- [Compute](#compute)
- [Deploy](#deploy)
- [Develop](#develop)
- [Distribute](#distribute)
- [Securing Artefacts](#securing-artefacts)
- [Securing Build Pipelines](#securing-build-pipelines)
- [Securing Deployments](#securing-deployments)
- [Securing Materials](#securing-materials)
- [Securing the Source Code](#securing-the-source-code)
- [Security Assurance](#security-assurance)
- [Storage](#storage)

---

## Access {#access}

Access

### Secrets are injected at runtime {#cnswp-1}

**Control ID**: `CNSWP-1`

#### Objective

Secrets are injected at runtime

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| IA-5(7) | 5 | Reference to IA-5(7) Authenticator Management |
| No Embedded Unencrypted Static Authenticators | 5 | Reference to No Embedded Unencrypted Static Authenticators |

---

### ABAC and RBAC are used {#cnswp-10}

**Control ID**: `CNSWP-10`

#### Objective

AC-3(7) Access Enforcement | Role-Based Access Control

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-3(13) | 5 | Reference to AC-3(13) Access Enforcement |
| Attribute-Based Access Control | 5 | Reference to Attribute-Based Access Control |

---

### End user identity is capable of being accepted, consumed, and forwarded on for contextual or dynamic authorization {#cnswp-11}

**Control ID**: `CNSWP-11`

#### Objective

End user identity is capable of being accepted, consumed, and forwarded on for contextual or dynamic authorization

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-7(19) | 5 | Reference to SC-7(19) Boundary Protection |
| Block Communication from Non-Organizationally Configured Hosts | 5 | Reference to Block Communication from Non-Organizationally Configured Hosts |

#### Recommendations

- This can be achieved through the use of identity documents and tokens.

---

### All cluster and workloads operators are authenticated {#cnswp-12}

**Control ID**: `CNSWP-12`

#### Objective

All cluster and workloads operators are authenticated

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| IA-7 Cryptographic Module Authentication | 5 | Reference to IA-7 Cryptographic Module Authentication |

---

### cluster and worklods operate actions are evaluated against access control policies governing context, purpose, and output {#cnswp-13}

**Control ID**: `CNSWP-13`

#### Objective

cluster and worklods operate actions are evaluated against access control policies governing context, purpose, and output

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| IA-7 Cryptographic Module Authentication | 5 | Reference to IA-7 Cryptographic Module Authentication |

---

### Identity federation uses multi-factor authentication {#cnswp-14}

**Control ID**: `CNSWP-14`

#### Objective

Identity federation uses multi-factor authentication

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| IA-2(1)(2) Identification and Authentication (organizational Users) | 5 | Reference to IA-2(1)(2) Identification and Authentication (organizational Users) |
| Multi-Factor Authenticaiton to Priviledged & Non Priveledged Accounts | 5 | Reference to Multi-Factor Authenticaiton to Priviledged & Non Priveledged Accounts |

---

### HSMs are used to physically protect cryptographic secrets with an encryption key residing in the HSM {#cnswp-15}

**Control ID**: `CNSWP-15`

#### Objective

SC-3(1) Security Function Isolation | Hardware Separation

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-4(4) Information Flow Enforcement | 5 | Reference to AC-4(4) Information Flow Enforcement |
| Flow Control of Encrypted Information | 5 | Reference to Flow Control of Encrypted Information |

#### Recommendations

- If this is not possible, software-based credential managers should be used.

---

### Secrets should have a short expiration period or time to live {#cnswp-16}

**Control ID**: `CNSWP-16`

#### Objective

Secrets should have a short expiration period or time to live

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-12 Information Management and Retention | 5 | Reference to SI-12 Information Management and Retention |

---

### time to live and expiration period on secrets is verfied to prevent reuse {#cnswp-17}

**Control ID**: `CNSWP-17`

#### Objective

time to live and expiration period on secrets is verfied to prevent reuse

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-16(3) Security and Privacy Attributes | 5 | Reference to AC-16(3) Security and Privacy Attributes |
| Maintenance of Attribute Associations by System | 5 | Reference to Maintenance of Attribute Associations by System |

---

### secrets management systems are highly available {#cnswp-18}

**Control ID**: `CNSWP-18`

#### Objective

secrets management systems are highly available

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-12(1) Cryptographic Key Establishment and Management | 5 | Reference to SC-12(1) Cryptographic Key Establishment and Management |
| Availability | 5 | Reference to Availability |

---

### long-lived secrets adhere to periodic rotation and revocation {#cnswp-19}

**Control ID**: `CNSWP-19`

#### Objective

long-lived secrets adhere to periodic rotation and revocation

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-12 Information Management and Retention | 5 | Reference to SI-12 Information Management and Retention |

#### Recommendations

- Long-lived secrets are not recommended, but some capabilities require them

---

### Applications and workloads are explicitly authorized to communicate with each other using mutual authentication {#cnswp-2}

**Control ID**: `CNSWP-2`

#### Objective

Applications and workloads are explicitly authorized to communicate with each other using mutual authentication

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| IA-9 | 5 | Reference to IA-9 Service Identification and Authentication |

---

### Secrets are distributed through secured communication channels protected commensurate with the level of access or data they are protecting {#cnswp-20}

**Control ID**: `CNSWP-20`

#### Objective

Secrets are distributed through secured communication channels protected commensurate with the level of access or data they are protecting

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-16 Security and Privacy Atributes | 5 | Reference to AC-16 Security and Privacy Atributes |

---

### Secrets injected are runtime are masqued or dropped from logs, audit, or system dumps {#cnswp-21}

**Control ID**: `CNSWP-21`

#### Objective

Secrets injected are runtime are masqued or dropped from logs, audit, or system dumps

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AU-9(3) Protection of Audit Information | 5 | Reference to AU-9(3) Protection of Audit Information |
| Cryptographic Protection | 5 | Reference to Cryptographic Protection |

#### Recommendations

- Even short lived secrets may be resused if caught in time by an interested attacker. Logs, audit, and systems dumps (i.e. in-memory shared volumes instead of environment variables) are all areas where runtime injected secrets show up

---

### Keys are rotated frequently {#cnswp-3}

**Control ID**: `CNSWP-3`

#### Objective

Keys are rotated frequently

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-12 | 5 | Reference to SC-12 Cryptographic Key Establishment and Management |

---

### Key lifespan is short {#cnswp-4}

**Control ID**: `CNSWP-4`

#### Objective

Key lifespan is short

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-12(3) | 5 | Reference to SC-12(3) Cryptographic Key Establishment and Management |
| Asymetric Key | 5 | Reference to Asymetric Key |

---

### Credentials and keys protecting sensitive workloads (health/finance/etc) are generated and managed independent of a cloud service provider {#cnswp-5}

**Control ID**: `CNSWP-5`

#### Objective

Credentials and keys protecting sensitive workloads (health/finance/etc) are generated and managed independent of a cloud service provider

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| IA-2(12) | 5 | Reference to IA-2(12) Identification and Authentication (Organizational Users) |
| Acceptance of PIV Credentials | 5 | Reference to Acceptance of PIV Credentials |

#### Recommendations

- KMS and HMS are common technologies to achive this. FIPS 140-2 complaince is strongly suggested. Cloud KMS tends to be FIPS 140-2 Level 2 or greater.

---

### Authentication and authorization are determined independently {#cnswp-6}

**Control ID**: `CNSWP-6`

#### Objective

Authentication and authorization are determined independently

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| IA-2(6) | 5 | Reference to IA-2(6) Identification and Authentication (Organizational Users) |
| Access to Accounts - Separate Devices | 5 | Reference to Access to Accounts - Separate Devices |

---

### Authentication and authorization are enforced independently {#cnswp-7}

**Control ID**: `CNSWP-7`

#### Objective

Authentication and authorization are enforced independently

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| IA-2(6) | 5 | Reference to IA-2(6) Identification and Authentication (Organizational Users) |
| Access to Accounts - Separate Devices | 5 | Reference to Access to Accounts - Separate Devices |

---

### access control and file permissions are updated in real-time {#cnswp-8}

**Control ID**: `CNSWP-8`

#### Objective

access control and file permissions are updated in real-time

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-4(2) | 5 | Reference to SI-4(2) System Monitoring |
| Automated Tools and Mechanisms for Real-Time Analysis | 5 | Reference to Automated Tools and Mechanisms for Real-Time Analysis |

#### Recommendations

- where possible as caching may permit unauthorized access

---

### authorization for workloads is granted based on attributs and roles/permissions previously assigned {#cnswp-9}

**Control ID**: `CNSWP-9`

#### Objective

authorization for workloads is granted based on attributs and roles/permissions previously assigned

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-3(13) | 5 | Reference to AC-3(13) Access Enforcement |
| Attribute-Based Access Control | 5 | Reference to Attribute-Based Access Control |

---

## Compute {#compute}

Compute

### Bootstrapping is employed to verify correct physical and logical location of compute {#cnswp-22}

**Control ID**: `CNSWP-22`

#### Objective

Bootstrapping is employed to verify correct physical and logical location of compute

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-7(9) Software, Firmware, and Information Integrity | 5 | Reference to SI-7(9) Software, Firmware, and Information Integrity |
| Verify Boot Process | 5 | Reference to Verify Boot Process |

#### Recommendations

- Secure Boot with TPM 2.0 or similar control

---

### Disparate data sensitive workloads are not run on the same OS kernel {#cnswp-23}

**Control ID**: `CNSWP-23`

#### Objective

Disparate data sensitive workloads are not run on the same OS kernel

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-7 Boundary Protection | 5 | Reference to SC-7 Boundary Protection |

#### Recommendations

- There are at least three implementing controls possible: workloads may be separated by running in a separate
cluster, on a separate node, or by implementing pods in independent VMs. It is also possible to emulate the kernel via
an application kernel (e.g. gvisor)


---

### Monitor and detect any changes to the initial configurations made in runtime {#cnswp-24}

**Control ID**: `CNSWP-24`

#### Objective

Monitor and detect any changes to the initial configurations made in runtime

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-2(2) Baseline Configuration | 5 | Reference to CM-2(2) Baseline Configuration |
| CM-3(7) | 5 |  |
| Review System Changes | 5 | Reference to Review System Changes |

#### Recommendations

- Preventative controls should be the primary control. Detective controls monitoring filesystem changes should be used to verify primary controls are operating properly.

---

### API auditing is enabled with a filter for a specific set of API Groups or verbs {#cnswp-25}

**Control ID**: `CNSWP-25`

#### Objective

API auditing is enabled with a filter for a specific set of API Groups or verbs

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AU-2 Event Logging | 5 | Reference to AU-2 Event Logging |

#### Recommendations

- API audits of the application, kubernetes API server, and kernel should be implemented.

---

### Container specific operating systems are in use {#cnswp-26}

**Control ID**: `CNSWP-26`

#### Objective

Container specific operating systems are in use

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-2CM-7 | 5 |  |

#### Recommendations

- a read-only OS with other services disabled. This provides isolation and resource confinement that enables developers to run isolated applications on a shared host kernel

---

### The hardware root of trust is based in a Trusted Platform Module (TPM) or virtual TPM (vTPM) {#cnswp-27}

**Control ID**: `CNSWP-27`

#### Objective

The hardware root of trust is based in a Trusted Platform Module (TPM) or virtual TPM (vTPM)

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-7 Software, Firmware, and Information Integrity | 5 | Reference to SI-7 Software, Firmware, and Information Integrity |

#### Recommendations

- Ensure HW root of trust extends to the OS kernel, modules, system images, container runtimes, and all software on the system.

---

### Minimize administrative access to the control plane {#cnswp-28}

**Control ID**: `CNSWP-28`

#### Objective

Minimize administrative access to the control plane

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-6 Least Privilege | 5 | Reference to AC-6 Least Privilege |

#### Recommendations

- Enure both users and pods have the minimum necessary access

---

### Object level and resource requests and limits are controlled through cgroups {#cnswp-29}

**Control ID**: `CNSWP-29`

#### Objective

SI-7(17)

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-7(16) Software, Firmware, and Information Integrity | 5 | Reference to SI-7(16) Software, Firmware, and Information Integrity |
| Time Limit on Process Execution Without Supervision | 5 | Reference to Time Limit on Process Execution Without Supervision |

#### Recommendations

- helps prevent exhaustion of node and cluster level resources by one misbehaving workload due to an intentional (e.g., fork bomb attack or cryptocurrency mining) or unintentional (e.g., reading a large file in memory without input validation, horizontal autoscaling to exhaust compute resources) issue

---

### Systems processing alerts are periodically tuned for false positives {#cnswp-30}

**Control ID**: `CNSWP-30`

#### Objective

Systems processing alerts are periodically tuned for false positives

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-4(13) System Monitoring | 5 | Reference to SI-4(13) System Monitoring |
| Analyze Traffic and Event Patterns | 5 | Reference to Analyze Traffic and Event Patterns |

#### Recommendations

- to avoid alert flooding, fatigue, and false negatives after security incidents that were not detected by the system

---

### All orchestrator control plane components are configured to communicate via mutual authentication and certificate validation with a periodically rotated certificate {#cnswp-31}

**Control ID**: `CNSWP-31`

#### Objective

All orchestrator control plane components are configured to communicate via mutual authentication and certificate validation with a periodically rotated certificate

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-3 Access Enforcement | 5 | Reference to AC-3 Access Enforcement |

#### Recommendations

- In unfederated clusters, the CA should be used exclusively for the current cluster.

---

### Only sanctioned capabilities and system calls (e.g. seccomp filters), are allowed to execute or be invoked in a container by the host operating system {#cnswp-32}

**Control ID**: `CNSWP-32`

#### Objective

Only sanctioned capabilities and system calls (e.g. seccomp filters), are allowed to execute or be invoked in a container by the host operating system

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-2,CM-7 | 5 |  |

#### Recommendations

- Additional tooling should be installed that go beyond k8s capabilities to limit system calls. E.g. Falco.

---

### Changes to critical mount points and files are prevented, monitored, and alerted {#cnswp-33}

**Control ID**: `CNSWP-33`

#### Objective

Changes to critical mount points and files are prevented, monitored, and alerted

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-5 Access Restrictions for Change | 5 | Reference to CM-5 Access Restrictions for Change |

---

### Runtime configuration control prevents changes to binaries, certificates, and remote access configurations {#cnswp-34}

**Control ID**: `CNSWP-34`

#### Objective

Runtime configuration control prevents changes to binaries, certificates, and remote access configurations

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-5 Access Restrictions for Change | 5 | Reference to CM-5 Access Restrictions for Change |

---

### Runtime configuration prevents ingress and egress network access for containers to only what is required to operate {#cnswp-35}

**Control ID**: `CNSWP-35`

#### Objective

Runtime configuration prevents ingress and egress network access for containers to only what is required to operate

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-7 Boundary Protection | 5 | Reference to SC-7 Boundary Protection |

---

### Policies are defined that restrict communications to only occur between sanctioned microservice pairs {#cnswp-36}

**Control ID**: `CNSWP-36`

#### Objective

Policies are defined that restrict communications to only occur between sanctioned microservice pairs

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-7 Boundary Protection | 5 | Reference to SC-7 Boundary Protection |

---

### Use a policy agent to control and enforce authorized, signed container images {#cnswp-37}

**Control ID**: `CNSWP-37`

#### Objective

Use a policy agent to control and enforce authorized, signed container images

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-5 Access Restrictions for Change | 5 | Reference to CM-5 Access Restrictions for Change |

---

### Use a policy agent to control provenance assurance for operational workloads {#cnswp-38}

**Control ID**: `CNSWP-38`

#### Objective

Use a policy agent to control provenance assurance for operational workloads

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-5 Access Restrictions for Change | 5 | Reference to CM-5 Access Restrictions for Change |

---

### Use a service mesh that eliminates implicit trust through data-in-motion encryption (data in transit) {#cnswp-39}

**Control ID**: `CNSWP-39`

#### Objective

Use a service mesh that eliminates implicit trust through data-in-motion encryption (data in transit)

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-7 Boundary Protection | 5 | Reference to SC-7 Boundary Protection |

---

### Use components that detect, track, aggregate and report system calls and network traffic from a container {#cnswp-40}

**Control ID**: `CNSWP-40`

#### Objective

Use components that detect, track, aggregate and report system calls and network traffic from a container

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-4 System Monitoring | 5 | Reference to SI-4 System Monitoring |

#### Recommendations

- should be leveraged to look for unexpected or malicious behavior

---

### Workloads should be dynamically scanned to detect malicious or insidious behavior for which no known occurrence yet exists {#cnswp-41}

**Control ID**: `CNSWP-41`

#### Objective

Workloads should be dynamically scanned to detect malicious or insidious behavior for which no known occurrence yet exists

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-3 Malicious Code Protection | 5 | Reference to SI-3 Malicious Code Protection |

#### Recommendations

- Events such as an extended sleep command that executes data exfiltration from etcd after the workload has been running for X amount of days are not expected in the majority of environments and therefore are not included in security tests. The aspect that workloads can have time or event delayed trojan horses is only detectable by comparing to baseline expected behavior, often discovered during thorough activity and scan monitoring

---

### Environments are continuously scanned to detect new vulnerabilities in workloads {#cnswp-42}

**Control ID**: `CNSWP-42`

#### Objective

Environments are continuously scanned to detect new vulnerabilities in workloads

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| RA-5 Vulnerability Monitoring and Scanning | 5 | Reference to RA-5 Vulnerability Monitoring and Scanning |

#### Recommendations

- Vulnerabilities are constantly being discovered, just because it wasnt vulnerable at deploy, doesn't mean it won't be vulnerable in two weeks

---

### Actionable audit events are generates that correlate/contextualize data from logs into "information" that can drive decision trees/incident response {#cnswp-43}

**Control ID**: `CNSWP-43`

#### Objective

Actionable audit events are generates that correlate/contextualize data from logs into "information" that can drive decision trees/incident response

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AU-3 Content of Audit Records | 5 | Reference to AU-3 Content of Audit Records |

---

### segregation of duties and the principle of least privilege is enforced {#cnswp-44}

**Control ID**: `CNSWP-44`

#### Objective

segregation of duties and the principle of least privilege is enforced

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-6 Least Privilege | 5 | Reference to AC-6 Least Privilege |

---

### Non-compliant violations are detected based on a pre-configured set of rules that filter violations of the organization's policies {#cnswp-45}

**Control ID**: `CNSWP-45`

#### Objective

Non-compliant violations are detected based on a pre-configured set of rules that filter violations of the organization's policies

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-7 Software, Firmware, and Information Integrity | 5 | Reference to SI-7 Software, Firmware, and Information Integrity |

---

### Native secret stores encrypt with keys from an external Key Management Store (KMS) {#cnswp-46}

**Control ID**: `CNSWP-46`

#### Objective

Native secret stores encrypt with keys from an external Key Management Store (KMS)

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-12(3) Systems & Communication Protection | 5 | Reference to SC-12(3) Systems & Communication Protection |

---

### Native secret stores are not configured for base64 encoding or stored in clear-text in the key-value store by default {#cnswp-47}

**Control ID**: `CNSWP-47`

#### Objective

Native secret stores are not configured for base64 encoding or stored in clear-text in the key-value store by default

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-12(3) Systems & Communication Protection | 5 | Reference to SC-12(3) Systems & Communication Protection |

#### Recommendations

- encoding is not encryption

---

### Network traffic to malicious domains is detected and denied {#cnswp-48}

**Control ID**: `CNSWP-48`

#### Objective

Network traffic to malicious domains is detected and denied

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-4 System Monitoring | 5 | Reference to SI-4 System Monitoring |

---

### Use encrypted containers for sensitive sources, methods, and data {#cnswp-49}

**Control ID**: `CNSWP-49`

#### Objective

Use encrypted containers for sensitive sources, methods, and data

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-28 Protection of Information at Rest | 5 | Reference to SC-28 Protection of Information at Rest |

---

### Use SBOMs to identify current deployments of vulnerable libraries, dependencies, and packages {#cnswp-50}

**Control ID**: `CNSWP-50`

#### Objective

Use SBOMs to identify current deployments of vulnerable libraries, dependencies, and packages

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-8 System Component Inventory | 5 | Reference to CM-8 System Component Inventory |

---

### Processes must execute only functions explicitly defined in an allow list {#cnswp-51}

**Control ID**: `CNSWP-51`

#### Objective

Processes must execute only functions explicitly defined in an allow list

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-2, CM-7 | 5 |  |

---

### Functions are not be allowed to make changes to critical file system mount points {#cnswp-52}

**Control ID**: `CNSWP-52`

#### Objective

Functions are not be allowed to make changes to critical file system mount points

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-5 Access Restrictions for Change | 5 | Reference to CM-5 Access Restrictions for Change |

---

### Function access is only permitted to sanctioned services {#cnswp-53}

**Control ID**: `CNSWP-53`

#### Objective

Function access is only permitted to sanctioned services

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-2, CM-7 | 5 |  |

#### Recommendations

- Either through networking restrictions or least privilege in permission models

---

### Egress network connection is monitored to detect and prevent access to C&C (command and control) and other malicious network domains {#cnswp-54}

**Control ID**: `CNSWP-54`

#### Objective

Egress network connection is monitored to detect and prevent access to C&C (command and control) and other malicious network domains

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-4 System Monitoring | 5 | Reference to SI-4 System Monitoring |

---

### Ingress network inspection is employed detect and remove malicious payloads and commands {#cnswp-55}

**Control ID**: `CNSWP-55`

#### Objective

Ingress network inspection is employed detect and remove malicious payloads and commands

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-4 System Monitoring | 5 | Reference to SI-4 System Monitoring |

#### Recommendations

- For instance, SQL injection attacks can be detected using inspection.

---

### Serverless functions are run in tenant-based resource or performance isolation for similar data classifications {#cnswp-56}

**Control ID**: `CNSWP-56`

#### Objective

Serverless functions are run in tenant-based resource or performance isolation for similar data classifications

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-7(21) Boundary Protection | 5 | Reference to SC-7(21) Boundary Protection |
| Isolation of System Components | 5 | Reference to Isolation of System Components |

#### Recommendations

- This may impact the performance due to limitations in the address space available to the isolation environment and should be considered for only the most sensitive workloads.

---

## Deploy {#deploy}

Deploy

### trust confirmation verifies the image has a valid signature from an authorized source {#cnswp-57}

**Control ID**: `CNSWP-57`

#### Objective

SR-4 (4)

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SR-4 (3) PROVENANCE | 5 | Reference to SR-4 (3) PROVENANCE |
| VALIDATE AS GENUINE AND NOT ALTERED | 5 | Reference to VALIDATE AS GENUINE AND NOT ALTERED |

---

### Image runtime policies are enforced prior to deployment {#cnswp-58}

**Control ID**: `CNSWP-58`

#### Objective

Image runtime policies are enforced prior to deployment

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-7 (17) SOFTWARE, FIRMWARE, AND INFORMATION INTEGRITY | 5 | Reference to SI-7 (17) SOFTWARE, FIRMWARE, AND INFORMATION INTEGRITY |
| RUNTIME APPLICATION SELF-PROTECTION | 5 | Reference to RUNTIME APPLICATION SELF-PROTECTION |

---

### Image integrity and signature are verifying prior to deployment {#cnswp-59}

**Control ID**: `CNSWP-59`

#### Objective

SR-4 (4)

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SR-4 (3) PROVENANCE | 5 | Reference to SR-4 (3) PROVENANCE |
| VALIDATE AS GENUINE AND NOT ALTERED | 5 | Reference to VALIDATE AS GENUINE AND NOT ALTERED |

---

### Applications provide logs regarding authentication, authorization, actions, and failures {#cnswp-60}

**Control ID**: `CNSWP-60`

#### Objective

Applications provide logs regarding authentication, authorization, actions, and failures

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-3 CONFIGURATION CHANGE CONTROL | 5 | Reference to CM-3 CONFIGURATION CHANGE CONTROL |

---

### Forensics capabilities are integrated into an incident response plan and procedures {#cnswp-61}

**Control ID**: `CNSWP-61`

#### Objective

Forensics capabilities are integrated into an incident response plan and procedures

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| INCIDENT HANDLING | 5 | Reference to INCIDENT HANDLING |
| MALICIOUS CODE AND FORENSIC ANALYSIS | 5 | Reference to MALICIOUS CODE AND FORENSIC ANALYSIS |

---

### AI, ML, or statistical modeling are used for behavioural and heuristic environment analysis {#cnswp-62}

**Control ID**: `CNSWP-62`

#### Objective

AI, ML, or statistical modeling are used for behavioural and heuristic environment analysis

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-3 SYSTEM AND INFORMATION INTEGRITY | 5 | Reference to SI-3 SYSTEM AND INFORMATION INTEGRITY |

---

## Develop {#develop}

Develop

### Establish a dedicated Production environment {#cnswp-63}

**Control ID**: `CNSWP-63`

#### Objective

Establish a dedicated Production environment

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-3(1) SYSTEM DEVELOPMENT LIFE CYCLE | 5 | Reference to SA-3(1) SYSTEM DEVELOPMENT LIFE CYCLE |
| MANAGE PREPRODUCTION ENVIRONMENT | 5 | Reference to MANAGE PREPRODUCTION ENVIRONMENT |

---

### Leverage Dynamic deployments {#cnswp-64}

**Control ID**: `CNSWP-64`

#### Objective

Leverage Dynamic deployments

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-8(31) SECURITY AND PRIVACY ENGINEERING PRINCIPLES | 5 | Reference to SA-8(31) SECURITY AND PRIVACY ENGINEERING PRINCIPLES |
| SECURE SYSTEM MODIFICATION | 5 | Reference to SECURE SYSTEM MODIFICATION |

#### Recommendations

- Blue/Green, Alpha/Beta, Canary, red-black deployments

---

### Integrate vulnerability and configuration scanning in the IDE or at the pull request {#cnswp-65}

**Control ID**: `CNSWP-65`

#### Objective

Integrate vulnerability and configuration scanning in the IDE or at the pull request

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-11(1) DEVELOPER TESTING AND EVALUATION | 5 | Reference to SA-11(1) DEVELOPER TESTING AND EVALUATION |
| STATIC CODE ANALYSIS | 5 | Reference to STATIC CODE ANALYSIS |

---

### Establish dedicated development, testing, and production environment {#cnswp-66}

**Control ID**: `CNSWP-66`

#### Objective

Establish dedicated development, testing, and production environment

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-15 DEVELOPMENT PROCESS, STANDARDS, AND TOOLS | 5 | Reference to SA-15 DEVELOPMENT PROCESS, STANDARDS, AND TOOLS |

---

### Build tests for business-critical code {#cnswp-67}

**Control ID**: `CNSWP-67`

#### Objective

Build tests for business-critical code

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-11 DEVELOPER TESTING AND EVALUATION | 5 | Reference to SA-11 DEVELOPER TESTING AND EVALUATION |

---

### Build tests for business-critical infrastructure {#cnswp-68}

**Control ID**: `CNSWP-68`

#### Objective

Build tests for business-critical infrastructure

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-11 DEVELOPER TESTING AND EVALUATION | 5 | Reference to SA-11 DEVELOPER TESTING AND EVALUATION |

---

### Test suite able to be ran locally {#cnswp-69}

**Control ID**: `CNSWP-69`

#### Objective

Test suite able to be ran locally

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-11 DEVELOPER TESTING AND EVALUATION | 5 | Reference to SA-11 DEVELOPER TESTING AND EVALUATION |

---

### Test suites should be available to run in a shared environment {#cnswp-70}

**Control ID**: `CNSWP-70`

#### Objective

Test suites should be available to run in a shared environment

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-11 DEVELOPER TESTING AND EVALUATION | 5 | Reference to SA-11 DEVELOPER TESTING AND EVALUATION |

---

### Implement two non-author reviewers/approvers prior to merging {#cnswp-71}

**Control ID**: `CNSWP-71`

#### Objective

Implement two non-author reviewers/approvers prior to merging

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-11(4) DEVELOPER TESTING AND EVALUATION | 5 | Reference to SA-11(4) DEVELOPER TESTING AND EVALUATION |
| MANUAL CODE REVIEWS | 5 | Reference to MANUAL CODE REVIEWS |

---

### Code should be clean and well commented {#cnswp-72}

**Control ID**: `CNSWP-72`

#### Objective

Code should be clean and well commented

---

### Full infrastructure tests are used {#cnswp-73}

**Control ID**: `CNSWP-73`

#### Objective

Full infrastructure tests are used

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-11 DEVELOPER TESTING AND EVALUATION | 5 | Reference to SA-11 DEVELOPER TESTING AND EVALUATION |

---

### Regression tests are used {#cnswp-74}

**Control ID**: `CNSWP-74`

#### Objective

Regression tests are used

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-11 DEVELOPER TESTING AND EVALUATION | 5 | Reference to SA-11 DEVELOPER TESTING AND EVALUATION |

---

### Test suites are updated against new and emerging threats and developed into security regressions tests {#cnswp-75}

**Control ID**: `CNSWP-75`

#### Objective

Test suites are updated against new and emerging threats and developed into security regressions tests

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-11 DEVELOPER TESTING AND EVALUATION | 5 | Reference to SA-11 DEVELOPER TESTING AND EVALUATION |

---

### Establish a dedicated Testing environment {#cnswp-76}

**Control ID**: `CNSWP-76`

#### Objective

Establish a dedicated Testing environment

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-3(1) SYSTEM DEVELOPMENT LIFE CYCLE | 5 | Reference to SA-3(1) SYSTEM DEVELOPMENT LIFE CYCLE |
| MANAGE PREPRODUCTION ENVIRONMENT | 5 | Reference to MANAGE PREPRODUCTION ENVIRONMENT |

---

### Continuous integration server is isolated {#cnswp-77}

**Control ID**: `CNSWP-77`

#### Objective

Continuous integration server is isolated

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-39 PROCESS ISOLATION | 5 | Reference to SC-39 PROCESS ISOLATION |

---

### Use threat model results to determine ROI for test development {#cnswp-78}

**Control ID**: `CNSWP-78`

#### Objective

Use threat model results to determine ROI for test development

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-11(2) DEVELOPER TESTING AND EVALUATION | 5 | Reference to SA-11(2) DEVELOPER TESTING AND EVALUATION |
| THREAT MODELING AND VULNERABILITY ANALYSES | 5 | Reference to THREAT MODELING AND VULNERABILITY ANALYSES |

---

## Distribute {#distribute}

Distribute

### Registries require mutually authenticated TLS for all registry connections {#cnswp-100}

**Control ID**: `CNSWP-100`

#### Objective

Registries require mutually authenticated TLS for all registry connections

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| IA-3(1) CRYPTOGRAPHIC BIDIRECTIONAL AUTHENTICATION | 5 | Reference to IA-3(1) CRYPTOGRAPHIC BIDIRECTIONAL AUTHENTICATION |

---

### image and metadata are signed {#cnswp-101}

**Control ID**: `CNSWP-101`

#### Objective

image and metadata are signed

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-7 SOFTWARE, FIRMWARE, AND INFORMATION INTEGRITY | 5 | Reference to SI-7 SOFTWARE, FIRMWARE, AND INFORMATION INTEGRITY |

---

### configuration is signed {#cnswp-102}

**Control ID**: `CNSWP-102`

#### Objective

configuration is signed

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-7 SOFTWARE, FIRMWARE, AND INFORMATION INTEGRITY | 5 | Reference to SI-7 SOFTWARE, FIRMWARE, AND INFORMATION INTEGRITY |

---

### package is signed {#cnswp-103}

**Control ID**: `CNSWP-103`

#### Objective

package is signed

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-7 SOFTWARE, FIRMWARE, AND INFORMATION INTEGRITY | 5 | Reference to SI-7 SOFTWARE, FIRMWARE, AND INFORMATION INTEGRITY |

---

### Validate integrity of images {#cnswp-104}

**Control ID**: `CNSWP-104`

#### Objective

Validate integrity of images

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-7 SYSTEM & INFORMATION INTEGRITY | 5 | Reference to SI-7 SYSTEM & INFORMATION INTEGRITY |

---

### Scan images for vulnerabilities and malware {#cnswp-105}

**Control ID**: `CNSWP-105`

#### Objective

SA-3

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| RA-5 VULNERABILITY MONITORING AND SCANNING | 5 | Reference to RA-5 VULNERABILITY MONITORING AND SCANNING |

---

### Enable image signing key revokation in the event of compromise {#cnswp-106}

**Control ID**: `CNSWP-106`

#### Objective

Enable image signing key revokation in the event of compromise

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-7 SYSTEM & INFORMATION INTEGRITY | 5 | Reference to SI-7 SYSTEM & INFORMATION INTEGRITY |

---

### Security updates are prioritized {#cnswp-107}

**Control ID**: `CNSWP-107`

#### Objective

Security updates are prioritized

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-2(3) SYSTEM & INFORMATION INTEGRITY | 5 | Reference to SI-2(3) SYSTEM & INFORMATION INTEGRITY |

---

### HSMs or credential managers should be used for protecting credentials {#cnswp-108}

**Control ID**: `CNSWP-108`

#### Objective

HSMs or credential managers should be used for protecting credentials

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-12(3) SYSTEMS & COMMUNICATION PROTECTION | 5 | Reference to SC-12(3) SYSTEMS & COMMUNICATION PROTECTION |

---

### Container image scanning findings are acted upon {#cnswp-109}

**Control ID**: `CNSWP-109`

#### Objective

Container image scanning findings are acted upon

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-2(3) SYSTEM & INFORMATION INTEGRITY | 5 | Reference to SI-2(3) SYSTEM & INFORMATION INTEGRITY |

---

### organizational compliance rules are enforced {#cnswp-110}

**Control ID**: `CNSWP-110`

#### Objective

organizational compliance rules are enforced

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| PL-1 POLICY AND PROCEDURES | 5 | Reference to PL-1 POLICY AND PROCEDURES |

---

### Incremental hardening of the infrastructure is employed {#cnswp-111}

**Control ID**: `CNSWP-111`

#### Objective

Incremental hardening of the infrastructure is employed

---

### pulls from public registries are controlled and only from authorized engineers or internal registries {#cnswp-112}

**Control ID**: `CNSWP-112`

#### Objective

pulls from public registries are controlled and only from authorized engineers or internal registries

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-6(3) LEAST PRIVILEGE | 5 | Reference to AC-6(3) LEAST PRIVILEGE |
| NETWORK ACCESS TO PRIVILEGED COMMANDS | 5 | Reference to NETWORK ACCESS TO PRIVILEGED COMMANDS |

---

### Image encryption is coupled with key management attestation and/or authorization and credential distribution {#cnswp-113}

**Control ID**: `CNSWP-113`

#### Objective

SC-12(3)

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-12(2) CRYPTOGRAPHIC KEY ESTABLISHMENT AND MANAGEMENT | 5 | Reference to SC-12(2) CRYPTOGRAPHIC KEY ESTABLISHMENT AND MANAGEMENT |
| SYMMETRIC & ASYMMETRIC KEYS | 5 | Reference to SYMMETRIC & ASYMMETRIC KEYS |

#### Recommendations

- useful for compliance use cases such as geo-fencing or export control and digital rights media management

---

### At-risk applications are prioritized for remediation by the exploit maturity and vulnerable path presence in addition to the CVSS score {#cnswp-114}

**Control ID**: `CNSWP-114`

#### Objective

At-risk applications are prioritized for remediation by the exploit maturity and vulnerable path presence in addition to the CVSS score

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-2(3) SYSTEM & INFORMATION INTEGRITY | 5 | Reference to SI-2(3) SYSTEM & INFORMATION INTEGRITY |

---

### Trust is verified {#cnswp-79}

**Control ID**: `CNSWP-79`

#### Objective

Trust is verified

---

### Artifacts ready for deployment are managed in a staging or pre-prod registry {#cnswp-80}

**Control ID**: `CNSWP-80`

#### Objective

Artifacts ready for deployment are managed in a staging or pre-prod registry

---

### container images are hardened following best practices {#cnswp-81}

**Control ID**: `CNSWP-81`

#### Objective

container images are hardened following best practices

#### Recommendations

- Images contain least permissions to remain functional, do not allow for shell, do not include unnecessary libraries and dependencies, do not bind mount files in from the host, etc.

---

### Static application security testing (SAST) is performed {#cnswp-82}

**Control ID**: `CNSWP-82`

#### Objective

Static application security testing (SAST) is performed

#### Recommendations

- Linting & fuzzing is performed

---

### Test suites follow the test pyramid {#cnswp-83}

**Control ID**: `CNSWP-83`

#### Objective

Test suites follow the test pyramid

---

### Artifacts undergoing active development are held in a private registery {#cnswp-84}

**Control ID**: `CNSWP-84`

#### Objective

Artifacts undergoing active development are held in a private registery

---

### Scan application manifests in CI pipeline {#cnswp-85}

**Control ID**: `CNSWP-85`

#### Objective

Scan application manifests in CI pipeline

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| RA-5 VULNERABILITY MONITORING AND SCANNING | 5 | Reference to RA-5 VULNERABILITY MONITORING AND SCANNING |

---

### CI server's for sensitive workloads are isolated from other workloads {#cnswp-86}

**Control ID**: `CNSWP-86`

#### Objective

CI server's for sensitive workloads are isolated from other workloads

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-39 PROCESS ISOLATION | 5 | Reference to SC-39 PROCESS ISOLATION |

---

### Builds requiring elevated privileges must run on dedicated servers {#cnswp-87}

**Control ID**: `CNSWP-87`

#### Objective

Builds requiring elevated privileges must run on dedicated servers

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-39 PROCESS ISOLATION | 5 | Reference to SC-39 PROCESS ISOLATION |

---

### Build policies are enforced on the CI pipeline {#cnswp-88}

**Control ID**: `CNSWP-88`

#### Objective

Build policies are enforced on the CI pipeline

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-1 POLICY AND PROCEDURES | 5 | Reference to SA-1 POLICY AND PROCEDURES |

---

### Sign pipeline metadata {#cnswp-89}

**Control ID**: `CNSWP-89`

#### Objective

Sign pipeline metadata

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-7 SOFTWARE, FIRMWARE, AND INFORMATION INTEGRITY | 5 | Reference to SI-7 SOFTWARE, FIRMWARE, AND INFORMATION INTEGRITY |

---

### Build stages are verified prior to the next stage executing {#cnswp-90}

**Control ID**: `CNSWP-90`

#### Objective

Build stages are verified prior to the next stage executing

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-7 SOFTWARE, FIRMWARE, AND INFORMATION INTEGRITY | 5 | Reference to SI-7 SOFTWARE, FIRMWARE, AND INFORMATION INTEGRITY |

---

### Images are scanned within the CI pipeline {#cnswp-91}

**Control ID**: `CNSWP-91`

#### Objective

SA-3

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| RA-5 VULNERABILITY MONITORING AND SCANNING | 5 | Reference to RA-5 VULNERABILITY MONITORING AND SCANNING |

---

### Vulnerability scans are coupled with pipeline compliance rules {#cnswp-92}

**Control ID**: `CNSWP-92`

#### Objective

Vulnerability scans are coupled with pipeline compliance rules

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-1 POLICY AND PROCEDURES | 5 | Reference to SA-1 POLICY AND PROCEDURES |

#### Recommendations

- Prevent insecure images and artifacts from being deployed

---

### Dynamic application security testing (DAST) is performed {#cnswp-93}

**Control ID**: `CNSWP-93`

#### Objective

Dynamic application security testing (DAST) is performed

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-11 (8) & (9) INTERACTIVE APPLICATION SECURITY TESTING | 5 | Reference to SA-11 (8) & (9) INTERACTIVE APPLICATION SECURITY TESTING |

#### Recommendations

- mocking

---

### Application instrumentation is employed {#cnswp-94}

**Control ID**: `CNSWP-94`

#### Objective

Application instrumentation is employed

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-4 SYSTEM MONITORING | 5 | Reference to SI-4 SYSTEM MONITORING |

---

### Automated test results map back to requirements {#cnswp-95}

**Control ID**: `CNSWP-95`

#### Objective

Automated test results map back to requirements

#### Recommendations

- Requirements include feature, function, security, and complaince

---

### Infrastructure security tests must be employed {#cnswp-96}

**Control ID**: `CNSWP-96`

#### Objective

Infrastructure security tests must be employed

#### Recommendations

- firewall rules open to the world, overprivileged Identity & Access Management (IAM) policies, unauthenticated endpoints, etc

---

### Tests to verify the security health are executed at time of build and at time of deploy {#cnswp-97}

**Control ID**: `CNSWP-97`

#### Objective

Tests to verify the security health are executed at time of build and at time of deploy

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-4 SYSTEM MONITORING | 5 | Reference to SI-4 SYSTEM MONITORING |

#### Recommendations

- to evaluate any changes or regressions that may have occurred throughout the lifecycle.

---

### IaC is subject to the same pipeline policy controls as application code {#cnswp-98}

**Control ID**: `CNSWP-98`

#### Objective

IaC is subject to the same pipeline policy controls as application code

---

### Security testing is automated {#cnswp-99}

**Control ID**: `CNSWP-99`

#### Objective

CA-8

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-11 DEVELOPER TESTING AND EVALUATION | 5 | Reference to SA-11 DEVELOPER TESTING AND EVALUATION |

---

## Securing Artefacts {#securing-artefacts}

Securing Artefacts

### Every step in the build process should be signed/attested for process integrity {#cnswp-141}

**Control ID**: `CNSWP-141`

#### Objective

SI-7

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-1 POLICY AND PROCEDURES | 5 | Reference to SI-1 POLICY AND PROCEDURES |

#### Recommendations

- should include these collective signatures and itself be signed to give integrity to the completed artefact and all its associated metadata.

---

### Every step in the build process should verify the previously generated signatures {#cnswp-142}

**Control ID**: `CNSWP-142`

#### Objective

SI-7

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-1 POLICY AND PROCEDURES | 5 | Reference to SI-1 POLICY AND PROCEDURES |

#### Recommendations

- artefacts should all be validated using the signatures generated by each step in its build process to ensure compliance

---

### Use a framework to manage signing of artefacts. {#cnswp-143}

**Control ID**: `CNSWP-143`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| IA-5 AUTHENTICATOR MANAGEMENT | 5 | Reference to IA-5 AUTHENTICATOR MANAGEMENT |

#### Recommendations

- from a single root to the individual teams or developers who sign artefacts. It uses additional metadata to allow clients to verify the freshness of content in a repository and protect against common attacks on update systems48. Clients can make use of public keys to verify the contents of the repository.

---

### Use a store to manage attestations {#cnswp-144}

**Control ID**: `CNSWP-144`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-4(6)	INFORMATION FLOW ENFORCEMENT | 5 | Reference to AC-4(6)	INFORMATION FLOW ENFORCEMENT |
| METADATA | 5 | Reference to METADATA |

#### Recommendations

- needs to be stored and tracked for which a database or a dedicated store such as Grafeas can be used.

---

### Limit which artefacts any given party is authorized to certify {#cnswp-145}

**Control ID**: `CNSWP-145`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-6 LEAST PRIVILEGE | 5 | Reference to AC-6 LEAST PRIVILEGE |

#### Recommendations

- to certify should be restricted using selective trust delegations. Trust must expire at predefined intervals, unless renewed as weel as a party must only be trusted to perform the tasks assigned to it to ensure compartmentatlization

---

### Rotation and revokation of private keys should be supported {#cnswp-146}

**Control ID**: `CNSWP-146`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-12 CRYPTOGRAPHIC KEY ESTABLISHMENT AND MANAGEMENT | 5 | Reference to SC-12 CRYPTOGRAPHIC KEY ESTABLISHMENT AND MANAGEMENT |

#### Recommendations

- and revoke private keys must be built into the distribution mechanism. Additionally, multiple keys must be used for different tasks or roles, and a threshold of keys must be required for important roles. Finally, minimal trust must be placed in high-risk keys like those that are stored online or used in automated roles.

---

### Use a container registry that supports OCI image-spec images {#cnswp-147}

**Control ID**: `CNSWP-147`

#### Recommendations

- with the security properties described in this section.

---

### Encrypt artefacts before distribution & ensure only authorized platforms have decryption capabilities {#cnswp-148}

**Control ID**: `CNSWP-148`

#### Objective

SC-13 CRYPTOGRAPHIC PROTECTION
SC-8 TRANSMISSION CONFIDENTIALITY AND INTEGRITY
IA-5 AUTHENTICATOR MANAGEMENT
SC-12 CRYPTOGRAPHIC KEY ESTABLISHMENT AND MANAGEMENT


#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-28(1)" | 5 |  |
| CRYPTOGRAPHIC PROTECTION | 5 | Reference to CRYPTOGRAPHIC PROTECTION |

#### Recommendations

- artefacts can be encrypted so that they are accessible by authorized parties, such as the clusters, vulnerability scanners, etc. t is recommended organizations use key management and distribution systems with identity and attestation mechanisms (e.g. SPIFFE/SPIRE)

---

## Securing Build Pipelines {#securing-build-pipelines}

Securing Build Pipelines

### Cryptographically guarantee policy adherence {#cnswp-149}

**Control ID**: `CNSWP-149`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-3(6)	CONFIGURATION CHANGE CONTROL | 5 | Reference to CM-3(6)	CONFIGURATION CHANGE CONTROL |
| CRYPTOGRAPHY MANAGEMENT | 5 | Reference to CRYPTOGRAPHY MANAGEMENT |

#### Recommendations

- in-toto project that can be used to secure a chain of pipeline stages end-to-end with cryptographic guarantees. Build metadata should be evaluated against the policy template by using tools such as Open Policy Agent.

---

### Validate environments and dependencies before usage {#cnswp-150}

**Control ID**: `CNSWP-150`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-3(2)	CONFIGURATION CHANGE CONTROL | 5 | Reference to CM-3(2)	CONFIGURATION CHANGE CONTROL |
| TESTING, VALIDATION, AND DOCUMENTATION OF CHANGES | 5 | Reference to TESTING, VALIDATION, AND DOCUMENTATION OF CHANGES |

#### Recommendations

- and any signatures should be validated both in the downloading or ingestion process, and again by the build worker. This should include validating package manager signatures, checking out specific Git commit hashes, and verifying SHA sums of input sources and binaries. After completing this validation, the downloading process should sign all binaries or libraries it is adding to the secure source

---

### Validate runtime security of build workers {#cnswp-151}

**Control ID**: `CNSWP-151`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-3(4) | 5 |  |
| SECURITY AND PRIVACY REPRESENTATIVES | 5 | Reference to SECURITY AND PRIVACY REPRESENTATIVES |

#### Recommendations

- such as seccomp, AppArmor, and SELinux, provides defense in depth against attacks on build infrastructure. High privilege kernel capabilities such as debugger, device, and network attachments should be restricted and monitored.

---

### Validate build artefacts through verifiably reproducible builds {#cnswp-152}

**Control ID**: `CNSWP-152`

#### Objective

Validate build artefacts through verifiably reproducible builds

CM-3(5)	CONFIGURATION CHANGE CONTROL | AUTOMATED SECURITY RESPONSE

Assurance Level: High

Risk Categories: High

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-3(4)" | 5 |  |
| SECURITY AND PRIVACY REPRESENTATIVES | 5 | Reference to SECURITY AND PRIVACY REPRESENTATIVES |

#### Recommendations

- build instructions, an end user should be able to reproduce the built artefact bit for bit.

---

### Lock and Verify External Requirements from the build process {#cnswp-153}

**Control ID**: `CNSWP-153`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-3(2) | 5 |  |
| TESTING, VALIDATION, AND DOCUMENTATION OF CHANGES | 5 | Reference to TESTING, VALIDATION, AND DOCUMENTATION OF CHANGES |

---

### Find and Eliminate Sources of Non-Determinism {#cnswp-154}

**Control ID**: `CNSWP-154`

#### Recommendations

- to dig in and find the cause of differences when tracking down sources of non-determinism.

---

### Record the Build Environment {#cnswp-155}

**Control ID**: `CNSWP-155`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-3(1) | 5 |  |
| AUTOMATED DOCUMENTATION, NOTIFICATION, AND PROHIBITION OF CHANGES | 5 | Reference to AUTOMATED DOCUMENTATION, NOTIFICATION, AND PROHIBITION OF CHANGES |

#### Recommendations

- layer

---

### Automate Creation of the Build Environment {#cnswp-156}

**Control ID**: `CNSWP-156`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-3(3) | 5 |  |
| AUTOMATED CHANGE IMPLEMENTATION | 5 | Reference to AUTOMATED CHANGE IMPLEMENTATION |

---

### Distribute Builds across different infrastructure {#cnswp-157}

**Control ID**: `CNSWP-157`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-3(3) | 5 |  |
| AUTOMATED CHANGE IMPLEMENTATION | 5 | Reference to AUTOMATED CHANGE IMPLEMENTATION |

---

### Build and related CI/CD steps should be automated through a pipeline delivered as code {#cnswp-158}

**Control ID**: `CNSWP-158`

#### Objective

SA-11

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-3 SYSTEM DEVELOPMENT LIFE CYCLE | 5 | Reference to SA-3 SYSTEM DEVELOPMENT LIFE CYCLE |

---

### Standardize pipelines across projects {#cnswp-159}

**Control ID**: `CNSWP-159`

---

### Provision a secured orchestration platform to host software factory {#cnswp-160}

**Control ID**: `CNSWP-160`

---

### Build workers should be single use {#cnswp-161}

**Control ID**: `CNSWP-161`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-2 ACCOUNT MANAGEMENT | 5 | Reference to AC-2 ACCOUNT MANAGEMENT |

---

### Ensure software factory has minimal network connectivity {#cnswp-162}

**Control ID**: `CNSWP-162`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-7(3)	BOUNDARY PROTECTION | 5 | Reference to SC-7(3)	BOUNDARY PROTECTION |
| ACCESS POINTS | 5 | Reference to ACCESS POINTS |

#### Recommendations

- of hardened source code, the dependency repository and code signing infrastructure.

---

### Segregate the duties of each build worker {#cnswp-163}

**Control ID**: `CNSWP-163`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-5 SEPARATION OF DUTIES | 5 | Reference to AC-5 SEPARATION OF DUTIES |

---

### Pass in build worker environment and commands {#cnswp-164}

**Control ID**: `CNSWP-164`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-2(2) BASELINE CONFIGURATION | 5 | Reference to CM-2(2) BASELINE CONFIGURATION |
| AUTOMATION SUPPORT FOR ACCURACY / CURRENC | 5 | Reference to AUTOMATION SUPPORT FOR ACCURACY / CURRENC |

#### Recommendations

- a clean and isolated environmment. It should not be able to pull its own environment. Ensure environment variables and commands are explicitly passed to avoid any complicated and opaque build process

---

### Write output to separate secured storage repo {#cnswp-165}

**Control ID**: `CNSWP-165`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AU-9(2) PROTECTION OF AUDIT INFORMATION | 5 | Reference to AU-9(2) PROTECTION OF AUDIT INFORMATION |
| STORE ON SEPARATE PHYSICAL SYSTEMS OR COMPONENTS | 5 | Reference to STORE ON SEPARATE PHYSICAL SYSTEMS OR COMPONENTS |

#### Recommendations

- from the Build Worker should then upload that artefact to an appropriate repository.

---

### Only allow pipeline modification through “pipeline as code” {#cnswp-166}

**Control ID**: `CNSWP-166`

#### Recommendations

- This prevents attackers from interacting and modifying the configuration. This model then requires appropriate authentication and authorization to be in place for the software and configuration of the pipeline

---

### Define user roles {#cnswp-167}

**Control ID**: `CNSWP-167`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-2 ACCOUNT MANAGEMENT | 5 | Reference to AC-2 ACCOUNT MANAGEMENT |

---

### Follow established practices for establishing a root of trust from an offline source {#cnswp-168}

**Control ID**: `CNSWP-168`

#### Objective

IA-5(2) AUTHENTICATOR MANAGEMENT | PUBLIC KEY-BASED AUTHENTICATION
SA-8(10) SECURITY AND PRIVACY ENGINEERING PRINCIPLES | HIERARCHICAL TRUST
SR-4(4)        PROVENANCE | SUPPLY CHAIN INTEGRITY — PEDIGREE


#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-17 PUBLIC KEY INFRASTRUCTURE CERTIFICATES | 5 | Reference to SC-17 PUBLIC KEY INFRASTRUCTURE CERTIFICATES |

---

### Use short-lived workload certificates {#cnswp-169}

**Control ID**: `CNSWP-169`

#### Objective

SC-17 PUBLIC KEY INFRASTRUCTURE CERTIFICATES


#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-23(5) SESSION AUTHENTICITY | 5 | Reference to SC-23(5) SESSION AUTHENTICITY |
| ALLOWED CERTIFICATE AUTHORITIES | 5 | Reference to ALLOWED CERTIFICATE AUTHORITIES |

---

## Securing Deployments {#securing-deployments}

Securing Deployments

### Ensure clients can perform verification of artefacts and associated metadata {#cnswp-170}

**Control ID**: `CNSWP-170`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-7 SOFTWARE, FIRMWARE, AND INFORMATION INTEGRITY | 5 | Reference to SI-7 SOFTWARE, FIRMWARE, AND INFORMATION INTEGRITY |

---

### Ensure clients can verify the “freshness” of files {#cnswp-171}

**Control ID**: `CNSWP-171`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-7 SOFTWARE, FIRMWARE, AND INFORMATION INTEGRITY | 5 | Reference to SI-7 SOFTWARE, FIRMWARE, AND INFORMATION INTEGRITY |

#### Recommendations

- Ensure clients can access latest versions and can veriify if the provided files are out of date

---

### Use a framework for managing software updates {#cnswp-172}

**Control ID**: `CNSWP-172`

#### Recommendations

- software updates in a secure, reliable and trusted way

---

## Securing Materials {#securing-materials}

Securing Materials

### Verify third party artefacts and open source libraries {#cnswp-173}

**Control ID**: `CNSWP-173`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-11 DEVELOPER TESTING AND EVALUATION | 5 | Reference to SA-11 DEVELOPER TESTING AND EVALUATION |

#### Recommendations

- of the continuous integration pipeline by validating their checksums against a known good source and validating any cryptographic signatures. Any software ingested must be scanned using Software Composition Analysis (SCA) and pentesting tools to detect whether any vulnerable open-source software is used in the final product.

---

### Require SBOM from third party suppliers {#cnswp-174}

**Control ID**: `CNSWP-174`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-8 INFORMATION SYSTEM COMPONENT INVENTORY | 5 | Reference to CM-8 INFORMATION SYSTEM COMPONENT INVENTORY |

#### Recommendations

- explicit details of the software and versions used within the supplied product as it provides a clear and direct link to the dependencies.

---

### Track dependencies between open source components {#cnswp-175}

**Control ID**: `CNSWP-175`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-10 SOFTWARE USAGE RESTRICTIONS | 5 | Reference to CM-10 SOFTWARE USAGE RESTRICTIONS |

#### Recommendations

- to help trace any deployed artefacts with new vulnerabilities. One of the most popular open source inventory implementations is OWASP Dependency-Track.

---

### Build libraries based upon source code {#cnswp-176}

**Control ID**: `CNSWP-176`

---

### Define and prioritize trusted package managers and repositories {#cnswp-177}

**Control ID**: `CNSWP-177`

#### Recommendations

- to pull from only those sources.

---

### Generate an immutable SBOM of the code {#cnswp-178}

**Control ID**: `CNSWP-178`

#### Recommendations

- There are currently two well known SBOM specifications: SPDX34 and CycloneDX

---

### Scan software for vulnerabilities {#cnswp-179}

**Control ID**: `CNSWP-179`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| RA-5 VULNERABILITY MONITORING AND SCANNING | 5 | Reference to RA-5 VULNERABILITY MONITORING AND SCANNING |

---

### Scan software for license implications {#cnswp-180}

**Control ID**: `CNSWP-180`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-10 SOFTWARE USAGE RESTRICTIONS | 5 | Reference to CM-10 SOFTWARE USAGE RESTRICTIONS |

#### Recommendations

- the Open Compliance Program36 which hosts several tools to ensure released software meets legal and regulatory compliance requirements.

---

### Run software composition analysis on ingested software {#cnswp-181}

**Control ID**: `CNSWP-181`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-11 (1) (8) & (9) DEVELOPER TESTING AND EVALUATION | 5 | Reference to SA-11 (1) (8) & (9) DEVELOPER TESTING AND EVALUATION |

#### Recommendations

- also serve as verification of SBOM content. This data will then be matched against data from a number of data feeds containing vulnerability data to highlight any vulnerabilities in the dependent packages.

---

## Securing the Source Code {#securing-the-source-code}

Securing the Source Code

### Commits and tags are signed {#cnswp-182}

**Control ID**: `CNSWP-182`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-7 SOFTWARE, FIRMWARE, AND INFORMATION INTEGRITY | 5 | Reference to SI-7 SOFTWARE, FIRMWARE, AND INFORMATION INTEGRITY |

#### Recommendations

- GPG keys or S/MIME certificates are used to sign the source code

---

### Enforce full attestation and verification for protected branches {#cnswp-183}

**Control ID**: `CNSWP-183`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-6(3) LEAST PRIVILEGE | 5 | Reference to AC-6(3) LEAST PRIVILEGE |
| NETWORK ACCESS TO PRIVILEGED COMMANDS | 5 | Reference to NETWORK ACCESS TO PRIVILEGED COMMANDS |

#### Recommendations

- Branch protection is enabled on the mainline and release branches with force push disabled

---

### Secrets are not committed to the source code repository unless encrypted {#cnswp-184}

**Control ID**: `CNSWP-184`

#### Objective

SC-12(2) CRYPTOGRAPHIC KEY ESTABLISHMENT AND MANAGEMENT | SYMMETRIC & ASYMMETRIC KEYS


#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-12(3) SYSTEMS & COMMUNICATION PROTECTION | 5 | Reference to SC-12(3) SYSTEMS & COMMUNICATION PROTECTION |

#### Recommendations

- Implement tooling to detect secrets or to prevent certain files from being pushed which may contain plaintext sensitive materials, such as via a .gitignore and/or .gitattributes file, client-side hook (pre-commit), server-side hook (pre-receive or update), and/or as a step in the CI process

---

### The individuals or teams with write access to a repository are defined {#cnswp-185}

**Control ID**: `CNSWP-185`

#### Objective

AC-3 ACCESS ENFORCEMENT


#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| PL-1 POLICY AND PROCEDURES | 5 | Reference to PL-1 POLICY AND PROCEDURES |

#### Recommendations

- Implement codeowners (or equivalent)

---

### Automate software security scanning and testing {#cnswp-186}

**Control ID**: `CNSWP-186`

#### Objective

SA-3 SYSTEM DEVELOPMENT LIFE CYCLE


#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| RA-5 VULNERABILITY MONITORING AND SCANNING | 5 | Reference to RA-5 VULNERABILITY MONITORING AND SCANNING |

#### Recommendations

- Security specific scans should be performed, including Static Application Security Tests (SAST) and Dynamic Application Security Tests (DAST). Both the coverage and results of these tests should be published as part of the repository information to help downstream consumers of software better assess the stability, reliability, and/or suitability of a product or library.

---

### Establish and adhere to contribution policies {#cnswp-187}

**Control ID**: `CNSWP-187`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| PL-1 POLICY AND PROCEDURES | 5 | Reference to PL-1 POLICY AND PROCEDURES |

#### Recommendations

- Define configuration options or configuration rules witthin SCM platforms allow repository administrators to enforce security, hygiene and operational policies.

---

### Define roles aligned to functional responsibilities {#cnswp-188}

**Control ID**: `CNSWP-188`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| PL-1 POLICY AND PROCEDURES | 5 | Reference to PL-1 POLICY AND PROCEDURES |

#### Recommendations

- Maintainer, Owner, Reviewer, Approver, and Guest

---

### Enforce an independent four-eyes principle {#cnswp-189}

**Control ID**: `CNSWP-189`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-11 DEVELOPER TESTING AND EVALUATION | 5 | Reference to SA-11 DEVELOPER TESTING AND EVALUATION |

#### Recommendations

- or greater expertise should review & approve the request.

---

### Use branch protection rules {#cnswp-190}

**Control ID**: `CNSWP-190`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-8 SECURITY ENGINEERING PRINCIPLES | 5 | Reference to SA-8 SECURITY ENGINEERING PRINCIPLES |

#### Recommendations

- Protection rules can be used to enforce the usage of pull requests with specified precondition and approval rules, ensuring that a human code review process is followed or an automated status checking of a branch occurs. Additionally, protected branches can be used to disallow dangerous use of force pushes26, preventing the overwrite of commit histories and potential obfuscation of code changes.

---

### Enforce MFA for accessing source code repositories {#cnswp-191}

**Control ID**: `CNSWP-191`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| IA-2(1) Identification and Authentication (organizational Users) | 5 | Reference to IA-2(1) Identification and Authentication (organizational Users) |
| Multi-Factor Authenticaiton to Priviledged Accounts | 5 | Reference to Multi-Factor Authenticaiton to Priviledged Accounts |

---

### Use SSH keys to provide developers access to source code repositories {#cnswp-192}

**Control ID**: `CNSWP-192`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-1 REMOTE ACCESS | 5 | Reference to AC-1 REMOTE ACCESS |

---

### Have a key rotation policy {#cnswp-193}

**Control ID**: `CNSWP-193`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-2(1) ACCOUNT MANAGEMENT | 5 | Reference to AC-2(1) ACCOUNT MANAGEMENT |
| AUTOMATED SYSTEM ACCOUNT MANAGEMENT | 5 | Reference to AUTOMATED SYSTEM ACCOUNT MANAGEMENT |

#### Recommendations

- usable after a certain period of time. When a private key is known to have been compromised, it should be revoked and replaced immediately to shut off access for any unauthorized user. Organizations may also consider using short lived certificates or keys, which reduces the reliance on certificate revocation systems.

---

### Use short-lived/ephemeral credentials for machine/service access {#cnswp-194}

**Control ID**: `CNSWP-194`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-2(1) ACCOUNT MANAGEMENT | 5 | Reference to AC-2(1) ACCOUNT MANAGEMENT |
| AUTOMATED SYSTEM ACCOUNT MANAGEMENT | 5 | Reference to AUTOMATED SYSTEM ACCOUNT MANAGEMENT |

#### Recommendations

- access tokens. For CI/CD pipeline agents, short-lived access tokens should be considered instead of password-based credentials. The use of very short-lived tokens like OAuth 2.0, OpenID Connect, etc., will help to implement more secure access and increase the security assurance.

---

## Security Assurance {#security-assurance}

Security Assurance

### Network policies enforce east-west network communication within the container deployment is limited to only that which is authorized for access {#cnswp-115}

**Control ID**: `CNSWP-115`

#### Objective

Network policies enforce east-west network communication within the container deployment is limited to only that which is authorized for access

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-6(3) LEAST PRIVILEGE | 5 | Reference to AC-6(3) LEAST PRIVILEGE |
| NETWORK ACCESS TO PRIVILEGED COMMANDS | 5 | Reference to NETWORK ACCESS TO PRIVILEGED COMMANDS |

---

### Incident reponse considers cloud native workloads {#cnswp-116}

**Control ID**: `CNSWP-116`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| IR-4 INCIDENT HANDLING | 5 | Reference to IR-4 INCIDENT HANDLING |
| AUTOMATED INCIDENT HANDLING PROCESSES | 5 | Reference to AUTOMATED INCIDENT HANDLING PROCESSES |

#### Recommendations

- instances could run on a different server), networking (e.g. IP addresses are assigned dynamically) and immutability (e.g. runtime changes to container are not persisted across restarts)

---

### Incident response accounts for appropriate evidence handling and collection of coud native workloads {#cnswp-117}

**Control ID**: `CNSWP-117`

#### Objective

Incident response accounts for appropriate evidence handling and collection of coud native workloads

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| IR-5(1) INCIDENT MONITORING | 5 | Reference to IR-5(1) INCIDENT MONITORING |
| AUTOMATED TRACKING, DATA COLLECTION, AND ANALYSIS | 5 | Reference to AUTOMATED TRACKING, DATA COLLECTION, AND ANALYSIS |

---

### Rootless builds are employed {#cnswp-118}

**Control ID**: `CNSWP-118`

#### Objective

Rootless builds are employed

---

### cgroups and system groups are used to isolate workloads and deployments {#cnswp-119}

**Control ID**: `CNSWP-119`

#### Objective

cgroups and system groups are used to isolate workloads and deployments

---

### MAC implementations are employed {#cnswp-120}

**Control ID**: `CNSWP-120`

#### Objective

MAC implementations are employed

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-3(3) ACCESS ENFORCEMENT | 5 | Reference to AC-3(3) ACCESS ENFORCEMENT |
| MANDATORY ACCESS CONTROL | 5 | Reference to MANDATORY ACCESS CONTROL |

#### Recommendations

- SELinux, AppArmor

---

### Threat model code and infrastructure {#cnswp-121}

**Control ID**: `CNSWP-121`

#### Objective

Threat model code and infrastructure

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-11(2) DEVELOPER TESTING AND EVALUATION | 5 | Reference to SA-11(2) DEVELOPER TESTING AND EVALUATION |
| THREAT MODELING AND VULNERABILITY ANALYSES | 5 | Reference to THREAT MODELING AND VULNERABILITY ANALYSES |

---

### Entities are able to independently authenticate other identities {#cnswp-122}

**Control ID**: `CNSWP-122`

#### Objective

Entities are able to independently authenticate other identities

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| IA-9 SERVICE IDENTIFICATION AND AUTHENTICATION | 5 | Reference to IA-9 SERVICE IDENTIFICATION AND AUTHENTICATION |

#### Recommendations

- Public Key Infrastructure

---

### Each entity can create proof of who the identity is {#cnswp-123}

**Control ID**: `CNSWP-123`

#### Objective

Each entity can create proof of who the identity is

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| IA-9 SERVICE IDENTIFICATION AND AUTHENTICATION | 5 | Reference to IA-9 SERVICE IDENTIFICATION AND AUTHENTICATION |

---

### Orchestrator is running on an a trusted OS, BIOS, etc {#cnswp-124}

**Control ID**: `CNSWP-124`

#### Objective

Orchestrator is running on an a trusted OS, BIOS, etc

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-14 SIGNED COMPONENTS | 5 | Reference to CM-14 SIGNED COMPONENTS |

---

### Orchestrator verifies the claims of a container {#cnswp-125}

**Control ID**: `CNSWP-125`

#### Objective

Orchestrator verifies the claims of a container

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-6 SECURITY AND PRIVACY FUNCTION VERIFICATION | 5 | Reference to SI-6 SECURITY AND PRIVACY FUNCTION VERIFICATION |

---

### Orchestrator network policies are used in conjunction with a service mesh {#cnswp-126}

**Control ID**: `CNSWP-126`

#### Objective

Orchestrator network policies are used in conjunction with a service mesh

---

## Storage {#storage}

Storage

### Storage control plane management interface requires mutual authentication and TLS for connections {#cnswp-127}

**Control ID**: `CNSWP-127`

#### Objective

Storage control plane management interface requires mutual authentication and TLS for connections

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-8 TRANSMISSION CONFIDENTIALITY AND INTEGRITY | 5 | Reference to SC-8 TRANSMISSION CONFIDENTIALITY AND INTEGRITY |

---

### Data availability is achieved through parity or mirroring, erasure coding or replicas {#cnswp-128}

**Control ID**: `CNSWP-128`

#### Objective

Data availability is achieved through parity or mirroring, erasure coding or replicas

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-13 PREDICTABLE FAILURE PREVENTION | 5 | Reference to SI-13 PREDICTABLE FAILURE PREVENTION |

---

### Hashing and checksums are added to blocks, objects or files {#cnswp-129}

**Control ID**: `CNSWP-129`

#### Objective

SI-7 SOFTWARE, FIRMWARE, AND INFORMATION INTEGRITY'


#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-7 LEAST FUNCTIONALITY | 5 | Reference to CM-7 LEAST FUNCTIONALITY |

#### Recommendations

- the tampering of data.

---

### Data backup storage locations employ like access controls and security policies to that of the data storage source {#cnswp-130}

**Control ID**: `CNSWP-130`

#### Objective

SC-30 CONCEALMENT AND MISDIRECTION'


#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-9 EXTERNAL SYSTEM SERVICES | 5 | Reference to SA-9 EXTERNAL SYSTEM SERVICES |

---

### Secure erasure adhering to OPAL standards is employed for returned or non-functional devices {#cnswp-131}

**Control ID**: `CNSWP-131`

#### Objective

MP-6 MEDIA SANITIZATION'


#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CP-9 SYSTEM BACKUP | 5 | Reference to CP-9 SYSTEM BACKUP |

---

### Encryption at rest considers data path, size, and frequency of access when determing additional security protections and cryptographic algorithms to employ {#cnswp-132}

**Control ID**: `CNSWP-132`

#### Objective

Encryption at rest considers data path, size, and frequency of access when determing additional security protections and cryptographic algorithms to employ

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-28 PROTECTION OF INFORMATION AT REST | 5 | Reference to SC-28 PROTECTION OF INFORMATION AT REST |

#### Recommendations

- will vary by system (e.g. per volume, per group or global keys)

---

### Caching is considered for determining encryption requirements in archictures {#cnswp-133}

**Control ID**: `CNSWP-133`

#### Objective

Caching is considered for determining encryption requirements in archictures

---

### Namespaces have defined trust boundaries to cordon access to volumes {#cnswp-134}

**Control ID**: `CNSWP-134`

#### Objective

Namespaces have defined trust boundaries to cordon access to volumes

---

### Security policies are used to prevent containers from accessing volume mounts on worker nodes {#cnswp-135}

**Control ID**: `CNSWP-135`

#### Objective

SA-8 SECURITY AND PRIVACY ENGINEERING PRINCIPLES
CM-6 CONFIGURATION SETTINGS'


#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-7 BOUNDARY PROTECTION | 5 | Reference to SC-7 BOUNDARY PROTECTION |

---

### Security policies are used enforce authorized worker node access to volumes {#cnswp-136}

**Control ID**: `CNSWP-136`

#### Objective

SA-8 SECURITY AND PRIVACY ENGINEERING PRINCIPLES
CM-6 CONFIGURATION SETTINGS'


#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-7 BOUNDARY PROTECTION | 5 | Reference to SC-7 BOUNDARY PROTECTION |

---

### Volume UID and GID are inaccessible to containers {#cnswp-137}

**Control ID**: `CNSWP-137`

#### Objective

AC-16 SECURITY AND PRIVACY ATTRIBUTES
SI-7 SOFTWARE, FIRMWARE, AND INFORMATION INTEGRITY'


#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-4 INFORMATION FLOW ENFORCEMENT | 5 | Reference to AC-4 INFORMATION FLOW ENFORCEMENT |

---

### Artifact registry supports OCI artifacts {#cnswp-138}

**Control ID**: `CNSWP-138`

#### Objective

Artifact registry supports OCI artifacts

---

### Artifact registry supports signed artifacts {#cnswp-139}

**Control ID**: `CNSWP-139`

#### Objective

Artifact registry supports signed artifacts

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-14 SIGNED COMPONENTS | 5 | Reference to CM-14 SIGNED COMPONENTS |

---

### Artifact registry verifies artifacts against organizational policies {#cnswp-140}

**Control ID**: `CNSWP-140`

#### Objective

Artifact registry verifies artifacts against organizational policies
CM-6 CONFIGURATION SETTINGS


#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AU-10 NON-REPUDIATION | 5 | Reference to AU-10 NON-REPUDIATION |

---

