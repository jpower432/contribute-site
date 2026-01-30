---
title: Cloud Native Security Controls Catalog
sidebar_position: 6
---

<!--
This file is auto-generated. Do not edit manually.

To regenerate this file, run below from the controls-catalog directory:
  go run cmd/catalog/main.go -md index.md
-->

# Cloud Native Security Controls Catalog

This catalog provides a structured framework for implementing security best practices in cloud-native environments.
It synthesizes the foundational principles of the [Cloud Native Security Whitepaper](../security-white%20paper)  and the Software Supply Chain Best Practices Paper into discrete, actionable objectives.

> **Note**: While this catalog is historically called the "Cloud Native Security Controls Catalog," these security objectives are expressed as **guidelines**. Throughout this document, we use "guidelines" to refer to the individual security recommendations, while "catalog" refers to the overall collection.

Guidelines are organized into Families, each representing a specific security domain. These families help you navigate and understand the scope of security guidelines across different aspects of cloud native systems.

Each entry contains the following components:
- **Guideline ID**: A unique identifier for traceability and mapping.
- **Objective**: The high-level security goal or intent of the guideline.
- **Guideline Mappings**: Cross-references to frameworks (e.g., NIST SP800-53r5) to support compliance alignment.
- **Statements**: Explanatory context and detailed descriptions of the guidelines requirements.
- **Recommendations**: Practical, non-binding guidance for implementation.

## Guideline Families

The following families organize guidelines by security domain. Click on any family name to jump to its guidelines:


### Access Control

Guidelines for access control models and identity forwarding. [View guidelines →](#cnswp-1)


### Compute

Guidelines for securing compute infrastructure including bootstrapping, isolation, monitoring, and runtime security. [View guidelines →](#cnswp-22)


### Deploy

Guidelines for securing software deployments, ensuring artifact verification, freshness validation, and secure update management. [View guidelines →](#cnswp-57)


### Distribute

Guidelines for secure distribution of container images, packages, and artifacts including signing, scanning, and registry security. [View guidelines →](#cnswp-100)


### Securing Artefacts

Guidelines for securing artefacts, including signing, verification, and freshness validation. [View guidelines →](#cnswp-141)


### Securing Build Pipelines

Guidelines for securing build pipelines, ensuring cryptographic guarantees, validation, and secure build environments. [View guidelines →](#cnswp-149)


### Securing Deployments

Guidelines for securing software deployments, ensuring artifact verification, freshness validation, and secure update management. [View guidelines →](#cnswp-170)


### Securing Materials

Guidelines for securing materials, including signing, verification, and freshness validation. [View guidelines →](#cnswp-173)


### Securing the Source Code

Guidelines for securing the source code, including signing, verification, and freshness validation. [View guidelines →](#cnswp-182)


### Security Assurance

Guidelines for security assurance, including signing, verification, and freshness validation. [View guidelines →](#cnswp-115)


### Storage

Guidelines for securing storage, including signing, verification, and freshness validation. [View guidelines →](#cnswp-127)



---


## Access Control {#access}

Guidelines for access control models and identity forwarding.


### Secrets are injected at runtime {#cnswp-1}

**Guideline ID**: `CNSWP-1`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| IA-5(7) |  | Authenticator Management |


#### Statements

Inject secrets at runtime rather than embedding them in code or configuration files.

---


### ABAC and RBAC are used {#cnswp-10}

**Guideline ID**: `CNSWP-10`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-3(13) |  | Access Enforcement |


#### Statements

Implement abac and rbac are used.

---


### End user identity is capable of being accepted, consumed, and forwarded on for contextual or dynamic authorization {#cnswp-11}

**Guideline ID**: `CNSWP-11`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-7(19) |  | Boundary Protection |


#### Statements

**Identity Forwarding Implementation**

Enable end user identity to be accepted, consumed, and forwarded.

---


### All cluster and workloads operators are authenticated {#cnswp-12}

**Guideline ID**: `CNSWP-12`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| IA-7 |  | Cryptographic Module Authentication |


#### Statements

Implement all cluster and workloads operators are authenticated.

---


### cluster and worklods operate actions are evaluated against access control policies governing context, purpose, and output {#cnswp-13}

**Guideline ID**: `CNSWP-13`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| IA-7 |  | Cryptographic Module Authentication |


#### Statements

Implement cluster and worklods operate actions are evaluated against access control policies governing context, purpose, and output.

---


### Identity federation uses multi-factor authentication {#cnswp-14}

**Guideline ID**: `CNSWP-14`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| IA-2(1)(2) |  | Identification and Authentication (organizational Users) |


#### Statements

Enforce multi-factor authentication for repository access.

---


### HSMs are used to physically protect cryptographic secrets with an encryption key residing in the HSM {#cnswp-15}

**Guideline ID**: `CNSWP-15`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-4(4) |  | Information Flow Enforcement |


#### Statements

**HSM Implementation**

Use HSMs to physically protect cryptographic secrets.

---


### Secrets should have a short expiration period or time to live {#cnswp-16}

**Guideline ID**: `CNSWP-16`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-12 |  | Information Management and Retention |


#### Statements

Configure secrets with short expiration periods or time-to-live values.

---


### time to live and expiration period on secrets is verified to prevent reuse {#cnswp-17}

**Guideline ID**: `CNSWP-17`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-16(3) |  | Security and Privacy Attributes |


#### Statements

Verify secret expiration and time-to-live values to prevent reuse of expired secrets.

---


### secrets management systems are highly available {#cnswp-18}

**Guideline ID**: `CNSWP-18`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-12(1) |  | Cryptographic Key Establishment and Management |


#### Statements

Implement secrets management systems are highly available.

---


### long-lived secrets adhere to periodic rotation and revocation {#cnswp-19}

**Guideline ID**: `CNSWP-19`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-12 |  | Information Management and Retention |


#### Statements

**Long-Lived Secret Management**

Manage long-lived secrets through periodic rotation and revocation.

---


### Applications and workloads are explicitly authorized to communicate with each other using mutual authentication {#cnswp-2}

**Guideline ID**: `CNSWP-2`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| IA-9 |  |  |


#### Statements

Use mutual authentication to verify the identity of both communicating parties.

---


### Secrets are distributed through secured communication channels protected commensurate with the level of access or data they are protecting {#cnswp-20}

**Guideline ID**: `CNSWP-20`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-16 |  | Security and Privacy Atributes |


#### Statements

Use secure communication channels for secret distribution with protection appropriate to the sensitivity level.

---


### Secrets injected are runtime are masqued or dropped from logs, audit, or system dumps {#cnswp-21}

**Guideline ID**: `CNSWP-21`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AU-9(3) |  | Protection of Audit Information |


#### Statements

**Secret Masking**

Mask or drop secrets from logs, audit, or system dumps.

---


### Keys are rotated frequently {#cnswp-3}

**Guideline ID**: `CNSWP-3`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-12 |  | Cryptographic Key Establishment and Management |


#### Statements

Implement keys are rotated frequently.

---


### Key lifespan is short {#cnswp-4}

**Guideline ID**: `CNSWP-4`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-12(3) |  | Cryptographic Key Establishment and Management |


#### Statements

Implement key lifespan is short.

---


### Credentials and keys protecting sensitive workloads (health/finance/etc) are generated and managed independent of a cloud service provider {#cnswp-5}

**Guideline ID**: `CNSWP-5`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| IA-2(12) |  | Identification and Authentication (Organizational Users) |


#### Statements

**Independent Key Management**

Generate and manage credentials and keys independent of cloud service providers.

---


### Authentication and authorization are determined independently {#cnswp-6}

**Guideline ID**: `CNSWP-6`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| IA-2(6) |  | Identification and Authentication (Organizational Users) |


#### Statements

Determine authentication and authorization independently.

---


### Authentication and authorization are enforced independently {#cnswp-7}

**Guideline ID**: `CNSWP-7`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| IA-2(6) |  | Identification and Authentication (Organizational Users) |


#### Statements

Enforce authentication and authorization independently.

---


### access control and file permissions are updated in real-time {#cnswp-8}

**Guideline ID**: `CNSWP-8`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-4(2) |  | System Monitoring |


#### Statements

**Real-Time Access Control**

Update access control and file permissions in real-time.

---


### authorization for workloads is granted based on attributs and roles/permissions previously assigned {#cnswp-9}

**Guideline ID**: `CNSWP-9`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-3(13) |  | Access Enforcement |


#### Statements

Implement authorization for workloads is granted based on attributs and roles/permissions previously assigned.

---




## Compute {#compute}

Guidelines for securing compute infrastructure including bootstrapping, isolation, monitoring, and runtime security.


### Bootstrapping is employed to verify correct physical and logical location of compute {#cnswp-22}

**Guideline ID**: `CNSWP-22`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-7(9) |  | Software, Firmware, and Information Integrity |


#### Statements

**Bootstrapping Implementation**

Implement bootstrapping to verify compute location and boot integrity.

---


### Disparate data sensitive workloads are not run on the same OS kernel {#cnswp-23}

**Guideline ID**: `CNSWP-23`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-7 |  | Boundary Protection |


#### Statements

**Workload Separation**

Separate disparate data sensitive workloads to prevent running on the same OS kernel.

---


### Monitor and detect any changes to the initial configurations made in runtime {#cnswp-24}

**Guideline ID**: `CNSWP-24`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-2(2) |  | Baseline Configuration |
| CM-3(7) |  |  |


#### Statements

**Configuration Monitoring**

Monitor and detect changes to initial configurations made at runtime.

---


### API auditing is enabled with a filter for a specific set of API Groups or verbs {#cnswp-25}

**Guideline ID**: `CNSWP-25`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AU-2 |  | Event Logging |


#### Statements

**API Auditing Implementation**

Enable API auditing with filters for specific API Groups or verbs.

---


### Container specific operating systems are in use {#cnswp-26}

**Guideline ID**: `CNSWP-26`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-2 |  |  |


#### Statements

**Container OS Implementation**

Use container-specific operating systems.

---


### The hardware root of trust is based in a Trusted Platform Module (TPM) or virtual TPM (vTPM) {#cnswp-27}

**Guideline ID**: `CNSWP-27`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-7 |  | Software, Firmware, and Information Integrity |


#### Statements

**Root of Trust Implementation**

Base hardware root of trust in TPM or vTPM.

---


### Minimize administrative access to the control plane {#cnswp-28}

**Guideline ID**: `CNSWP-28`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-6 |  | Least Privilege |


#### Statements

**Access Minimization**

Minimize administrative access to the control plane.

---


### Object level and resource requests and limits are controlled through cgroups {#cnswp-29}

**Guideline ID**: `CNSWP-29`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-7(16) |  | Software, Firmware, and Information Integrity |


#### Statements

**Resource Control Implementation**

Control object level and resource requests and limits through cgroups.

---


### Systems processing alerts are periodically tuned for false positives {#cnswp-30}

**Guideline ID**: `CNSWP-30`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-4(13) |  | System Monitoring |


#### Statements

**Alert Tuning**

Periodically tune alert systems to reduce false positives.

---


### All orchestrator control plane components are configured to communicate via mutual authentication and certificate validation with a periodically rotated certificate {#cnswp-31}

**Guideline ID**: `CNSWP-31`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-3 |  | Access Enforcement |


#### Statements

**Control Plane Communication**

Configure all orchestrator control plane components for secure communication.

---


### Only sanctioned capabilities and system calls (e.g. seccomp filters), are allowed to execute or be invoked in a container by the host operating system {#cnswp-32}

**Guideline ID**: `CNSWP-32`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-2 |  |  |


#### Statements

**System Call Restriction**

Restrict capabilities and system calls allowed in containers.

---


### Changes to critical mount points and files are prevented, monitored, and alerted {#cnswp-33}

**Guideline ID**: `CNSWP-33`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-5 |  | Access Restrictions for Change |


#### Statements

Implement changes to critical mount points and files are prevented, monitored, and alerted.

---


### Runtime configuration control prevents changes to binaries, certificates, and remote access configurations {#cnswp-34}

**Guideline ID**: `CNSWP-34`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-5 |  | Access Restrictions for Change |


#### Statements

Implement runtime configuration control prevents changes to binaries, certificates, and remote access configurations.

---


### Runtime configuration prevents ingress and egress network access for containers to only what is required to operate {#cnswp-35}

**Guideline ID**: `CNSWP-35`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-7 |  | Boundary Protection |


#### Statements

Implement runtime configuration prevents ingress and egress network access for containers to only what is required to operate.

---


### Policies are defined that restrict communications to only occur between sanctioned microservice pairs {#cnswp-36}

**Guideline ID**: `CNSWP-36`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-7 |  | Boundary Protection |


#### Statements

Implement policies are defined that restrict communications to only occur between sanctioned microservice pairs.

---


### Use a policy agent to control and enforce authorized, signed container images {#cnswp-37}

**Guideline ID**: `CNSWP-37`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-5 |  | Access Restrictions for Change |


#### Statements

Sign images and artifacts to ensure integrity and authenticity.

---


### Use a policy agent to control provenance assurance for operational workloads {#cnswp-38}

**Guideline ID**: `CNSWP-38`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-5 |  | Access Restrictions for Change |


#### Statements

Implement use a policy agent to control provenance assurance for operational workloads.

---


### Use a service mesh that eliminates implicit trust through data-in-motion encryption (data in transit) {#cnswp-39}

**Guideline ID**: `CNSWP-39`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-7 |  | Boundary Protection |


#### Statements

Implement use a service mesh that eliminates implicit trust through data-in-motion encryption (data in transit).

---


### Use components that detect, track, aggregate and report system calls and network traffic from a container {#cnswp-40}

**Guideline ID**: `CNSWP-40`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-4 |  | System Monitoring |


#### Statements

**Monitoring Implementation**

Use components to detect, track, aggregate and report system calls and network traffic.

---


### Workloads should be dynamically scanned to detect malicious or insidious behavior for which no known occurrence yet exists {#cnswp-41}

**Guideline ID**: `CNSWP-41`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-3 |  | Malicious Code Protection |


#### Statements

**Dynamic Scanning**

Dynamically scan workloads to detect malicious or insidious behavior.

---


### Environments are continuously scanned to detect new vulnerabilities in workloads {#cnswp-42}

**Guideline ID**: `CNSWP-42`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| RA-5 |  | Vulnerability Monitoring and Scanning |


#### Statements

**Continuous Scanning**

Continuously scan environments to detect new vulnerabilities.

---


### Actionable audit events are generates that correlate/contextualize data from logs into "information" that can drive decision trees/incident response {#cnswp-43}

**Guideline ID**: `CNSWP-43`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AU-3 |  | Content of Audit Records |


#### Statements

Implement actionable audit events are generates that correlate/contextualize data from logs into "information" that can drive decision trees/incident response.

---


### segregation of duties and the principle of least privilege is enforced {#cnswp-44}

**Guideline ID**: `CNSWP-44`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-6 |  | Least Privilege |


#### Statements

Implement segregation of duties and the principle of least privilege is enforced.

---


### Non-compliant violations are detected based on a pre-configured set of rules that filter violations of the organization's policies {#cnswp-45}

**Guideline ID**: `CNSWP-45`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-7 |  | Software, Firmware, and Information Integrity |


#### Statements

Implement non-compliant violations are detected based on a pre-configured set of rules that filter violations of the organization's policies.

---


### Native secret stores encrypt with keys from an external Key Management Store (KMS) {#cnswp-46}

**Guideline ID**: `CNSWP-46`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-12(3) |  | Systems & Communication Protection |


#### Statements

Implement native secret stores encrypt with keys from an external key management store (kms).

---


### Native secret stores are not configured for base64 encoding or stored in clear-text in the key-value store by default {#cnswp-47}

**Guideline ID**: `CNSWP-47`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-12(3) |  | Systems & Communication Protection |


#### Statements

**Secret Storage**

Ensure native secret stores do not use base64 encoding or clear-text storage.

---


### Network traffic to malicious domains is detected and denied {#cnswp-48}

**Guideline ID**: `CNSWP-48`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-4 |  | System Monitoring |


#### Statements

Implement network traffic to malicious domains is detected and denied.

---


### Use encrypted containers for sensitive sources, methods, and data {#cnswp-49}

**Guideline ID**: `CNSWP-49`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-28 |  | Protection of Information at Rest |


#### Statements

Implement use encrypted containers for sensitive sources, methods, and data.

---


### Use SBOMs to identify current deployments of vulnerable libraries, dependencies, and packages {#cnswp-50}

**Guideline ID**: `CNSWP-50`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-8 |  | System Component Inventory |


#### Statements

Implement use sboms to identify current deployments of vulnerable libraries, dependencies, and packages.

---


### Processes must execute only functions explicitly defined in an allow list {#cnswp-51}

**Guideline ID**: `CNSWP-51`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-2 |  |  |


#### Statements

Implement processes must execute only functions explicitly defined in an allow list.

---


### Functions are not be allowed to make changes to critical file system mount points {#cnswp-52}

**Guideline ID**: `CNSWP-52`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-5 |  | Access Restrictions for Change |


#### Statements

Implement functions are not be allowed to make changes to critical file system mount points.

---


### Function access is only permitted to sanctioned services {#cnswp-53}

**Guideline ID**: `CNSWP-53`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-2 |  |  |


#### Statements

**Service Access Restriction**

Restrict function access to only sanctioned services.

---


### Egress network connection is monitored to detect and prevent access to C&C (command and control) and other malicious network domains {#cnswp-54}

**Guideline ID**: `CNSWP-54`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-4 |  | System Monitoring |


#### Statements

Implement egress network connection is monitored to detect and prevent access to c&c (command and control) and other malicious network domains.

---


### Ingress network inspection is employed detect and remove malicious payloads and commands {#cnswp-55}

**Guideline ID**: `CNSWP-55`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-4 |  | System Monitoring |


#### Statements

**Network Inspection**

Employ ingress network inspection to detect and remove malicious payloads.

---


### Serverless functions are run in tenant-based resource or performance isolation for similar data classifications {#cnswp-56}

**Guideline ID**: `CNSWP-56`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-7(21) |  | Boundary Protection |


#### Statements

**Serverless Isolation**

Run serverless functions in tenant-based resource or performance isolation.

---




## Deploy {#deploy}

Guidelines for securing software deployments, ensuring artifact verification, freshness validation, and secure update management.


### trust confirmation verifies the image has a valid signature from an authorized source {#cnswp-57}

**Guideline ID**: `CNSWP-57`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SR-4(3) |  | (3) Provenance |


#### Statements

Sign images and artifacts to ensure integrity and authenticity.

---


### Image runtime policies are enforced prior to deployment {#cnswp-58}

**Guideline ID**: `CNSWP-58`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-7(17) |  | (17) Software, Firmware, and Information Integrity |


#### Statements

Implement image runtime policies are enforced prior to deployment.

---


### Image integrity and signature are verifying prior to deployment {#cnswp-59}

**Guideline ID**: `CNSWP-59`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SR-4(3) |  | (3) Provenance |


#### Statements

Sign images and artifacts to ensure integrity and authenticity.

---


### Applications provide logs regarding authentication, authorization, actions, and failures {#cnswp-60}

**Guideline ID**: `CNSWP-60`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-3 |  | Configuration Change Control |


#### Statements

Implement applications provide logs regarding authentication, authorization, actions, and failures.

---


### Forensics capabilities are integrated into an incident response plan and procedures {#cnswp-61}

**Guideline ID**: `CNSWP-61`

#### Statements

Integrate forensics capabilities into incident response plans and procedures.

---


### AI, ML, or statistical modeling are used for behavioural and heuristic environment analysis {#cnswp-62}

**Guideline ID**: `CNSWP-62`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-3 |  | System and Information Integrity |


#### Statements

Implement ai, ml, or statistical modeling are used for behavioural and heuristic environment analysis.

---




## Distribute {#distribute}

Guidelines for secure distribution of container images, packages, and artifacts including signing, scanning, and registry security.


### Registries require mutually authenticated TLS for all registry connections {#cnswp-100}

**Guideline ID**: `CNSWP-100`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| IA-3(1) |  | Cryptographic Bidirectional Authentication |


#### Statements

Require mutually authenticated TLS for all registry connections.

---


### image and metadata are signed {#cnswp-101}

**Guideline ID**: `CNSWP-101`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-7 |  | Software, Firmware, and Information Integrity |


#### Statements

Sign images and artifacts to ensure integrity and authenticity.

---


### configuration is signed {#cnswp-102}

**Guideline ID**: `CNSWP-102`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-7 |  | Software, Firmware, and Information Integrity |


#### Statements

Sign configuration files to ensure integrity.

---


### package is signed {#cnswp-103}

**Guideline ID**: `CNSWP-103`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-7 |  | Software, Firmware, and Information Integrity |


#### Statements

Sign packages to verify integrity and authenticity.

---


### Validate integrity of images {#cnswp-104}

**Guideline ID**: `CNSWP-104`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-7 |  | System and Information Integrity |


#### Statements

Implement validate integrity of images.

---


### Scan images for vulnerabilities and malware {#cnswp-105}

**Guideline ID**: `CNSWP-105`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| RA-5 |  | Vulnerability Monitoring and Scanning |


#### Statements

Scan artifacts for vulnerabilities before distribution or deployment.

---


### Enable image signing key revokation in the event of compromise {#cnswp-106}

**Guideline ID**: `CNSWP-106`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-7 |  | System and Information Integrity |


#### Statements

Implement enable image signing key revokation in the event of compromise.

---


### Security updates are prioritized {#cnswp-107}

**Guideline ID**: `CNSWP-107`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-2(3) |  | System and Information Integrity |


#### Statements

Implement security updates are prioritized.

---


### HSMs or credential managers should be used for protecting credentials {#cnswp-108}

**Guideline ID**: `CNSWP-108`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-12(3) |  | Systems and Communication Protection |


#### Statements

Use hardware security modules to physically protect cryptographic secrets.

---


### Container image scanning findings are acted upon {#cnswp-109}

**Guideline ID**: `CNSWP-109`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-2(3) |  | System and Information Integrity |


#### Statements

Scan artifacts for vulnerabilities before distribution or deployment.

---


### organizational compliance rules are enforced {#cnswp-110}

**Guideline ID**: `CNSWP-110`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| PL-1 |  | Policy and Procedures |


#### Statements

Implement organizational compliance rules are enforced.

---


### Incremental hardening of the infrastructure is employed {#cnswp-111}

**Guideline ID**: `CNSWP-111`

---


### pulls from public registries are controlled and only from authorized engineers or internal registries {#cnswp-112}

**Guideline ID**: `CNSWP-112`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-6(3) |  | Least Privilege |


#### Statements

Implement pulls from public registries are controlled and only from authorized engineers or internal registries.

---


### Image encryption is coupled with key management attestation and/or authorization and credential distribution {#cnswp-113}

**Guideline ID**: `CNSWP-113`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-12(2) |  | Cryptographic Key Establishment and Management |


#### Statements

**Encryption and Key Management**

Couple image encryption with key management, attestation, and authorization.

---


### At-risk applications are prioritized for remediation by the exploit maturity and vulnerable path presence in addition to the CVSS score {#cnswp-114}

**Guideline ID**: `CNSWP-114`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-2(3) |  | System and Information Integrity |


#### Statements

Implement at-risk applications are prioritized for remediation by the exploit maturity and vulnerable path presence in addition to the cvss score.

---


### Trust is verified {#cnswp-79}

**Guideline ID**: `CNSWP-79`

---


### Artifacts ready for deployment are managed in a staging or pre-prod registry {#cnswp-80}

**Guideline ID**: `CNSWP-80`

---


### container images are hardened following best practices {#cnswp-81}

**Guideline ID**: `CNSWP-81`

#### Statements

**Image Hardening Practices**

Harden container images following security best practices.

---


### Static application security testing (SAST) is performed {#cnswp-82}

**Guideline ID**: `CNSWP-82`

#### Statements

**SAST Implementation**

Perform static application security testing.

---


### Test suites follow the test pyramid {#cnswp-83}

**Guideline ID**: `CNSWP-83`

---


### Artifacts undergoing active development are held in a private registery {#cnswp-84}

**Guideline ID**: `CNSWP-84`

---


### Scan application manifests in CI pipeline {#cnswp-85}

**Guideline ID**: `CNSWP-85`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| RA-5 |  | Vulnerability Monitoring and Scanning |


#### Statements

Scan artifacts for vulnerabilities before distribution or deployment.

---


### CI server's for sensitive workloads are isolated from other workloads {#cnswp-86}

**Guideline ID**: `CNSWP-86`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-39 |  | Process Isolation |


#### Statements

Implement ci server's for sensitive workloads are isolated from other workloads.

---


### Builds requiring elevated privileges must run on dedicated servers {#cnswp-87}

**Guideline ID**: `CNSWP-87`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-39 |  | Process Isolation |


#### Statements

Implement builds requiring elevated privileges must run on dedicated servers.

---


### Build policies are enforced on the CI pipeline {#cnswp-88}

**Guideline ID**: `CNSWP-88`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-1 |  | Policy and Procedures |


#### Statements

Implement build policies are enforced on the ci pipeline.

---


### Sign pipeline metadata {#cnswp-89}

**Guideline ID**: `CNSWP-89`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-7 |  | Software, Firmware, and Information Integrity |


#### Statements

Sign pipeline metadata to ensure integrity.

---


### Build stages are verified prior to the next stage executing {#cnswp-90}

**Guideline ID**: `CNSWP-90`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-7 |  | Software, Firmware, and Information Integrity |


#### Statements

Implement build stages are verified prior to the next stage executing.

---


### Images are scanned within the CI pipeline {#cnswp-91}

**Guideline ID**: `CNSWP-91`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| RA-5 |  | Vulnerability Monitoring and Scanning |


#### Statements

Scan artifacts for vulnerabilities before distribution or deployment.

---


### Vulnerability scans are coupled with pipeline compliance rules {#cnswp-92}

**Guideline ID**: `CNSWP-92`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-1 |  | Policy and Procedures |


#### Statements

**Compliance Integration**

Couple vulnerability scans with pipeline compliance rules.

---


### Dynamic application security testing (DAST) is performed {#cnswp-93}

**Guideline ID**: `CNSWP-93`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-11(8) |  | (8) & (9) Interactive Application Security Testing |


#### Statements

**DAST Implementation**

Perform dynamic application security testing.

---


### Application instrumentation is employed {#cnswp-94}

**Guideline ID**: `CNSWP-94`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-4 |  | System Monitoring |


#### Statements

Implement application instrumentation is employed.

---


### Automated test results map back to requirements {#cnswp-95}

**Guideline ID**: `CNSWP-95`

#### Statements

**Test Mapping**

Map automated test results back to requirements.

---


### Infrastructure security tests must be employed {#cnswp-96}

**Guideline ID**: `CNSWP-96`

#### Statements

**Infrastructure Testing**

Employ infrastructure security tests.

---


### Tests to verify the security health are executed at time of build and at time of deploy {#cnswp-97}

**Guideline ID**: `CNSWP-97`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-4 |  | System Monitoring |


#### Statements

**Security Health Testing**

Execute security health tests at build and deploy time.

---


### IaC is subject to the same pipeline policy controls as application code {#cnswp-98}

**Guideline ID**: `CNSWP-98`

---


### Security testing is automated {#cnswp-99}

**Guideline ID**: `CNSWP-99`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-11 |  | Developer Testing and Evaluation |


#### Statements

Implement security testing is automated.

---




## Securing Artefacts {#securing-artefacts}

Guidelines for securing artefacts, including signing, verification, and freshness validation.


### Every step in the build process should be signed/attested for process integrity {#cnswp-141}

**Guideline ID**: `CNSWP-141`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-1 |  | Policy and Procedures |


#### Statements

**Build Step Signing**

Every step in the build process should be signed and attested, with collective signatures included and the completed artifact itself signed.

---


### Every step in the build process should verify the previously generated signatures {#cnswp-142}

**Guideline ID**: `CNSWP-142`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-1 |  | Policy and Procedures |


#### Statements

**Signature Verification Process**

Every step in the build process should verify previously generated signatures to ensure compliance.

---


### Use a framework to manage signing of artefacts {#cnswp-143}

**Guideline ID**: `CNSWP-143`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| IA-5 |  | Authenticator Management |


#### Statements

**Signing Framework Implementation**

Use a framework to manage signing of artifacts from a single root to individual teams or developers.

---


### Use a store to manage attestations {#cnswp-144}

**Guideline ID**: `CNSWP-144`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-4(6) |  | Information Flow Enforcement |


#### Statements

**Attestation Store Implementation**

Use a dedicated store to manage and track attestations.

---


### Limit which artefacts any given party is authorized to certify {#cnswp-145}

**Guideline ID**: `CNSWP-145`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-6 |  | Least Privilege |


#### Statements

**Certification Authorization**

Limit which artifacts any given party is authorized to certify using selective trust delegations.

---


### Rotation and revokation of private keys should be supported {#cnswp-146}

**Guideline ID**: `CNSWP-146`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-12 |  | Cryptographic Key Establishment and Management |


#### Statements

**Key Rotation and Revocation**

Support rotation and revocation of private keys in the distribution mechanism.

---


### Use a container registry that supports OCI image-spec images {#cnswp-147}

**Guideline ID**: `CNSWP-147`

#### Statements

**OCI Registry Selection**

Use a container registry that supports OCI image-spec images with the security properties described in this section.

---


### Encrypt artefacts before distribution & ensure only authorized platforms have decryption capabilities {#cnswp-148}

**Guideline ID**: `CNSWP-148`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-28(1) |  |  |


#### Statements

**Artifact Encryption**

Encrypt artifacts before distribution so they are accessible only by authorized parties.

---




## Securing Build Pipelines {#securing-build-pipelines}

Guidelines for securing build pipelines, ensuring cryptographic guarantees, validation, and secure build environments.


### Cryptographically guarantee policy adherence {#cnswp-149}

**Guideline ID**: `CNSWP-149`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-3(6) |  | Configuration Change Control |


#### Statements

**Cryptographic Policy Implementation**

Use frameworks like in-toto to secure a chain of pipeline stages end-to-end with cryptographic guarantees.

---


### Validate environments and dependencies before usage {#cnswp-150}

**Guideline ID**: `CNSWP-150`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-3(2) |  | Configuration Change Control |


#### Statements

**Validation Process**

Validate environments and dependencies, including signatures, both in the downloading or ingestion process, and again by the build worker.

---


### Validate runtime security of build workers {#cnswp-151}

**Guideline ID**: `CNSWP-151`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-3(4) |  |  |


#### Statements

**Runtime Security Implementation**

Use security mechanisms such as seccomp, AppArmor, and SELinux to provide defense in depth.

---


### Validate build artefacts through verifiably reproducible builds {#cnswp-152}

**Guideline ID**: `CNSWP-152`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-3(4) |  |  |


#### Statements

**Reproducible Build Requirements**

With build instructions, an end user should be able to reproduce the built artefact bit for bit.

---


### Lock and Verify External Requirements from the build process {#cnswp-153}

**Guideline ID**: `CNSWP-153`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-3(2) |  |  |


#### Statements

Implement lock and verify external requirements from the build process.

---


### Find and Eliminate Sources of Non-Determinism {#cnswp-154}

**Guideline ID**: `CNSWP-154`

#### Statements

**Non-Determinism Elimination**

Investigate and eliminate sources of non-determinism in build processes.

---


### Record the Build Environment {#cnswp-155}

**Guideline ID**: `CNSWP-155`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-3(1) |  |  |


#### Statements

**Environment Recording**

Record all aspects of the build environment including base layers.

---


### Automate Creation of the Build Environment {#cnswp-156}

**Guideline ID**: `CNSWP-156`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-3(3) |  |  |


#### Statements

Implement automate creation of the build environment.

---


### Distribute Builds across different infrastructure {#cnswp-157}

**Guideline ID**: `CNSWP-157`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-3(3) |  |  |


#### Statements

Implement distribute builds across different infrastructure.

---


### Build and related CI/CD steps should be automated through a pipeline delivered as code {#cnswp-158}

**Guideline ID**: `CNSWP-158`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-3 |  | System Development Life Cycle |


#### Statements

Implement build and related ci/cd steps should be automated through a pipeline delivered as code.

---


### Standardize pipelines across projects {#cnswp-159}

**Guideline ID**: `CNSWP-159`

---


### Provision a secured orchestration platform to host software factory {#cnswp-160}

**Guideline ID**: `CNSWP-160`

---


### Build workers should be single use {#cnswp-161}

**Guideline ID**: `CNSWP-161`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-2 |  | Account Management |


#### Statements

Implement build workers should be single use.

---


### Ensure software factory has minimal network connectivity {#cnswp-162}

**Guideline ID**: `CNSWP-162`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-7(3) |  | Boundary Protection |


#### Statements

**Network Isolation**

Ensure software factory has minimal network connectivity to only essential services.

---


### Segregate the duties of each build worker {#cnswp-163}

**Guideline ID**: `CNSWP-163`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-5 |  | Separation of Duties |


#### Statements

Implement segregate the duties of each build worker.

---


### Pass in build worker environment and commands {#cnswp-164}

**Guideline ID**: `CNSWP-164`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-2(2) |  | Baseline Configuration |


#### Statements

**Environment and Command Management**

Build workers should operate in a clean and isolated environment and not be able to pull their own environment.

---


### Write output to separate secured storage repo {#cnswp-165}

**Guideline ID**: `CNSWP-165`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AU-9(2) |  | Protection of Audit Information |


#### Statements

**Artifact Storage**

Build workers should upload artifacts to appropriate secured repositories.

---


### Only allow pipeline modification through "pipeline as code" {#cnswp-166}

**Guideline ID**: `CNSWP-166`

#### Statements

**Pipeline as Code Enforcement**

Only allow pipeline modification through pipeline as code to prevent attackers from interacting and modifying the configuration.

---


### Define user roles {#cnswp-167}

**Guideline ID**: `CNSWP-167`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-2 |  | Account Management |


#### Statements

Implement define user roles.

---


### Follow established practices for establishing a root of trust from an offline source {#cnswp-168}

**Guideline ID**: `CNSWP-168`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-17 |  | Public Key Infrastructure Certificates |


#### Statements

Implement bootstrapping to verify compute location and boot integrity.

---


### Use short-lived workload certificates {#cnswp-169}

**Guideline ID**: `CNSWP-169`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-23(5) |  | Session Authenticity |


#### Statements

Implement use short-lived workload certificates.

---




## Securing Deployments {#securing-deployments}

Guidelines for securing software deployments, ensuring artifact verification, freshness validation, and secure update management.


### Ensure clients can perform verification of artefacts and associated metadata {#cnswp-170}

**Guideline ID**: `CNSWP-170`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-7 |  | Software, Firmware, and Information Integrity |


#### Statements

**Verification Capabilities**

Provide mechanisms for clients to verify artifacts and associated metadata.

---


### Ensure clients can verify the "freshness" of files {#cnswp-171}

**Guideline ID**: `CNSWP-171`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-7 |  | Software, Firmware, and Information Integrity |


#### Statements

**Freshness Verification**

Ensure clients can access latest versions and can verify if the provided files are out of date.

---


### Use a framework for managing software updates {#cnswp-172}

**Guideline ID**: `CNSWP-172`

#### Statements

**Update Framework Implementation**

Use a framework for managing software updates in a secure, reliable and trusted way.

---




## Securing Materials {#securing-materials}

Guidelines for securing materials, including signing, verification, and freshness validation.


### Verify third party artefacts and open source libraries {#cnswp-173}

**Guideline ID**: `CNSWP-173`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-11 |  | Developer Testing and Evaluation |


#### Statements

**Verification Process**

Verify third-party artifacts and open source libraries as part of the continuous integration pipeline by validating their checksums against a known good source and validating any cryptographic signatures.

---


### Require SBOM from third party suppliers {#cnswp-174}

**Guideline ID**: `CNSWP-174`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-8 |  | Information System Component Inventory |


#### Statements

**SBOM Requirements**

Require SBOM from third-party suppliers with explicit details of the software and versions used within the supplied product.

---


### Track dependencies between open source components {#cnswp-175}

**Guideline ID**: `CNSWP-175`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-10 |  | Software Usage Restrictions |


#### Statements

**Dependency Tracking Implementation**

Track dependencies between open source components to help trace any deployed artifacts with new vulnerabilities.

---


### Build libraries based upon source code {#cnswp-176}

**Guideline ID**: `CNSWP-176`

#### Statements

**Source-Based Building**

Build libraries from source code to ensure integrity and enable verification of the build process.

---


### Define and prioritize trusted package managers and repositories {#cnswp-177}

**Guideline ID**: `CNSWP-177`

#### Statements

**Trusted Sources Management**

Define and prioritize trusted package managers and repositories to pull from only those sources.

---


### Generate an immutable SBOM of the code {#cnswp-178}

**Guideline ID**: `CNSWP-178`

#### Statements

**SBOM Generation**

Generate an immutable SBOM of the code for all software artifacts.

---


### Scan software for vulnerabilities {#cnswp-179}

**Guideline ID**: `CNSWP-179`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| RA-5 |  | Vulnerability Monitoring and Scanning |


#### Statements

**Vulnerability Scanning Process**

Scan software for vulnerabilities as part of the development and build process.

---


### Scan software for license implications {#cnswp-180}

**Guideline ID**: `CNSWP-180`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-10 |  | Software Usage Restrictions |


#### Statements

**License Scanning Implementation**

Scan software for license implications to ensure released software meets legal and regulatory compliance requirements.

---


### Run software composition analysis on ingested software {#cnswp-181}

**Guideline ID**: `CNSWP-181`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-11(1)(8) |  | (1) (8) & (9) Developer Testing and Evaluation |


#### Statements

**Software Composition Analysis Process**

Run software composition analysis on ingested software to identify components and verify SBOM content.

---




## Securing the Source Code {#securing-the-source-code}

Guidelines for securing the source code, including signing, verification, and freshness validation.


### Commits and tags are signed {#cnswp-182}

**Guideline ID**: `CNSWP-182`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-7 |  | Software, Firmware, and information integrity |


#### Statements

**Signing Implementation**

GPG keys or S/MIME certificates are used to sign the source code.

---


### Enforce full attestation and verification for protected branches {#cnswp-183}

**Guideline ID**: `CNSWP-183`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-6(3) |  | Least Privilege |


#### Statements

**Branch Protection Configuration**

Branch protection is enabled on the mainline and release branches with force push disabled.

---


### Secrets are not committed to the source code repository unless encrypted {#cnswp-184}

**Guideline ID**: `CNSWP-184`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-12(3) |  | Systems and Communication Protection |


#### Statements

**Secret Detection and Prevention**

Implement tooling to detect secrets or to prevent certain files from being pushed which may contain plaintext sensitive materials.

---


### The individuals or teams with write access to a repository are defined {#cnswp-185}

**Guideline ID**: `CNSWP-185`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| PL-1 |  | Policy and Procedures |


#### Statements

**Access Definition**

Implement CODEOWNERS (or equivalent) to define who has write access and responsibility for different parts of the codebase.

---


### Automate software security scanning and testing {#cnswp-186}

**Guideline ID**: `CNSWP-186`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| RA-5 |  | Vulnerability Monitoring and Scanning |


#### Statements

**Security Scanning Implementation**

Security specific scans should be performed, including Static Application Security Tests (SAST) and Dynamic Application Security Tests (DAST).

---


### Establish and adhere to contribution policies {#cnswp-187}

**Guideline ID**: `CNSWP-187`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| PL-1 |  | Policy and Procedures |


#### Statements

**Policy Definition**

Define configuration options or configuration rules within SCM platforms allow repository administrators to enforce security, hygiene and operational policies.

---


### Define roles aligned to functional responsibilities {#cnswp-188}

**Guideline ID**: `CNSWP-188`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| PL-1 |  | Policy and Procedures |


#### Statements

**Role Definitions**

Define and document roles with specific responsibilities and access levels.

---


### Enforce an independent four-eyes principle {#cnswp-189}

**Guideline ID**: `CNSWP-189`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-11 |  | Developer Testing and Evaluation |


#### Statements

**Review Requirements**

Require independent review and approval by reviewers with equal or greater expertise.

---


### Use branch protection rules {#cnswp-190}

**Guideline ID**: `CNSWP-190`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-8 |  | Security Engineering Principles |


#### Statements

**Branch Protection Configuration**

Protection rules can be used to enforce the usage of pull requests with specified precondition and approval rules.

---


### Enforce MFA for accessing source code repositories {#cnswp-191}

**Guideline ID**: `CNSWP-191`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| IA-2(1) |  | Identification and Authentication (organizational Users) |


#### Statements

Implement enforce mfa for accessing source code repositories.

---


### Use SSH keys to provide developers access to source code repositories {#cnswp-192}

**Guideline ID**: `CNSWP-192`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-1 |  | Remote Access |


#### Statements

Implement use ssh keys to provide developers access to source code repositories.

---


### Have a key rotation policy {#cnswp-193}

**Guideline ID**: `CNSWP-193`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-2(1) |  | Prerequisites and criteria for group and role membership are defined. |


#### Statements

**Key Rotation Requirements**

Keys should be rotated after a certain period of time or when compromised.

---


### Use short-lived/ephemeral credentials for machine/service access {#cnswp-194}

**Guideline ID**: `CNSWP-194`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-2(1) |  | Usage of automated mechanisms to create, enable, modify, disable, and remove accounts. |


#### Statements

**Ephemeral Credential Implementation**

Use short-lived access tokens for CI/CD pipeline agents and service access.

---




## Security Assurance {#security-assurance}

Guidelines for security assurance, including signing, verification, and freshness validation.


### Network policies enforce east-west network communication within the container deployment is limited to only that which is authorized for access {#cnswp-115}

**Guideline ID**: `CNSWP-115`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-6(3) |  | Least Privilege |


#### Statements

Implement network policies enforce east-west network communication within the container deployment is limited to only that which is authorized for access.

---


### Incident reponse considers cloud native workloads {#cnswp-116}

**Guideline ID**: `CNSWP-116`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| IR-4 |  | Incident Handling |


#### Statements

**Cloud Native Considerations**

Incident response must consider cloud native workload characteristics.

---


### Incident response accounts for appropriate evidence handling and collection of coud native workloads {#cnswp-117}

**Guideline ID**: `CNSWP-117`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| IR-5(1) |  | Incident Monitoring |


#### Statements

Implement incident response accounts for appropriate evidence handling and collection of coud native workloads.

---


### Rootless builds are employed {#cnswp-118}

**Guideline ID**: `CNSWP-118`

---


### cgroups and system groups are used to isolate workloads and deployments {#cnswp-119}

**Guideline ID**: `CNSWP-119`

---


### MAC implementations are employed {#cnswp-120}

**Guideline ID**: `CNSWP-120`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-3(3) |  | Access Enforcement |


#### Statements

**MAC Implementation**

Use Mandatory Access Control (MAC) implementations for workload security.

---


### Threat model code and infrastructure {#cnswp-121}

**Guideline ID**: `CNSWP-121`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-11(2) |  | Developer Testing and Evaluation |


#### Statements

Implement threat model code and infrastructure.

---


### Entities are able to independently authenticate other identities {#cnswp-122}

**Guideline ID**: `CNSWP-122`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| IA-9 |  | Service Identification and Authentication |


#### Statements

**Independent Authentication**

Entities should be able to independently authenticate other identities.

---


### Each entity can create proof of who the identity is {#cnswp-123}

**Guideline ID**: `CNSWP-123`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| IA-9 |  | Service Identification and Authentication |


#### Statements

Implement each entity can create proof of who the identity is.

---


### Orchestrator is running on an a trusted OS, BIOS, etc {#cnswp-124}

**Guideline ID**: `CNSWP-124`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-14 |  | Signed Components |


#### Statements

Implement orchestrator is running on an a trusted os, bios, etc.

---


### Orchestrator verifies the claims of a container {#cnswp-125}

**Guideline ID**: `CNSWP-125`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-6 |  | Security and Privacy Function Verification |


#### Statements

Implement orchestrator verifies the claims of a container.

---


### Orchestrator network policies are used in conjunction with a service mesh {#cnswp-126}

**Guideline ID**: `CNSWP-126`

---




## Storage {#storage}

Guidelines for securing storage, including signing, verification, and freshness validation.


### Storage control plane management interface requires mutual authentication and TLS for connections {#cnswp-127}

**Guideline ID**: `CNSWP-127`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-8 |  | Transmission Confidentiality and Integrity |


#### Statements

Use mutual authentication to verify the identity of both communicating parties.

---


### Data availability is achieved through parity or mirroring, erasure coding or replicas {#cnswp-128}

**Guideline ID**: `CNSWP-128`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SI-13 |  | Predictable Failure Prevention |


#### Statements

Implement data availability is achieved through parity or mirroring, erasure coding or replicas.

---


### Hashing and checksums are added to blocks, objects or files {#cnswp-129}

**Guideline ID**: `CNSWP-129`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-7 |  | Least Functionality |


#### Statements

**Integrity Protection**

Add hashing and checksums to blocks, objects or files to detect tampering of data.

---


### Data backup storage locations employ like access controls and security policies to that of the data storage source {#cnswp-130}

**Guideline ID**: `CNSWP-130`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SA-9 |  | External System Services |


#### Statements

Implement data backup storage locations employ like access controls and security policies to that of the data storage source.

---


### Secure erasure adhering to OPAL standards is employed for returned or non-functional devices {#cnswp-131}

**Guideline ID**: `CNSWP-131`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CP-9 |  | System Backup |


#### Statements

Implement secure erasure adhering to opal standards is employed for returned or non-functional devices.

---


### Encryption at rest considers data path, size, and frequency of access when determing additional security protections and cryptographic algorithms to employ {#cnswp-132}

**Guideline ID**: `CNSWP-132`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-28 |  | Protection of Information at Rest |


#### Statements

**Encryption Strategy**

Consider data path, size, and frequency of access when determining encryption approach.

---


### Caching is considered for determining encryption requirements in archictures {#cnswp-133}

**Guideline ID**: `CNSWP-133`

---


### Namespaces have defined trust boundaries to cordon access to volumes {#cnswp-134}

**Guideline ID**: `CNSWP-134`

---


### Security policies are used to prevent containers from accessing volume mounts on worker nodes {#cnswp-135}

**Guideline ID**: `CNSWP-135`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-7 |  | Boundary Protection |


#### Statements

Implement security policies are used to prevent containers from accessing volume mounts on worker nodes.

---


### Security policies are used enforce authorized worker node access to volumes {#cnswp-136}

**Guideline ID**: `CNSWP-136`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| SC-7 |  | Boundary Protection |


#### Statements

Implement security policies are used enforce authorized worker node access to volumes.

---


### Volume UID and GID are inaccessible to containers {#cnswp-137}

**Guideline ID**: `CNSWP-137`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AC-4 |  | Information Flow Enforcement |


#### Statements

Implement volume uid and gid are inaccessible to containers.

---


### Artifact registry supports OCI artifacts {#cnswp-138}

**Guideline ID**: `CNSWP-138`

---


### Artifact registry supports signed artifacts {#cnswp-139}

**Guideline ID**: `CNSWP-139`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| CM-14 |  | Signed Components |


#### Statements

Sign images and artifacts to ensure integrity and authenticity.

---


### Artifact registry verifies artifacts against organizational policies {#cnswp-140}

**Guideline ID**: `CNSWP-140`

#### Guideline Mappings

**NIST-800-53**

| Reference ID | Strength | Remarks |
|--------------|----------|----------|
| AU-10 |  | Non-repudiation |


#### Statements

Implement artifact registry verifies artifacts against organizational policies.

---





## Acknowledgements
This representation of the catalog builds upon the original [Cloud Native Security Controls Catalog initiative](https://www.cncf.io/blog/2022/06/07/introduction-to-the-cloud-native-security-controls-catalog/), which produced the foundational artifact.

This catalog is expressed in **Gemara Layer 1** (Guidance Document) format, where security objectives are represented as guidelines. See [Gemara Documentation](https://gemara.openssf.org/) for details.
