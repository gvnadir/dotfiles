# SailPoint - IdentityIQ

## Introducing IdentityIQ

SailPoint IdentityIQ provides:

- Operational efficiency
- Integration with existing systems
- Access to all your identities
- Compliance to enterprises with complex environments

### Benefits

- Stregthen security and lower risk
- Improve compliance and audit performance
- Deliver fast, efficient access for the business

### IdentityIQ Capabilities

**IdentityIQ Compliance Manager** and **Lifecycle Manager** modules provide access management for the applications used by an enterprise and they share:

- Data and Database (Application Access Data)
- User Interface
- Implementation process

**Integrated solutions** instead have separate:

- Databases
- User Interface
- Implementation processes
- AI-Driven Identity Security (discover, manage, and control all your access-related data using machine learning and data intelligence)
- File Access Manager (helps locate and secure sensitive data, meet data privacy requirements and certify data access)

#### Lifecycle Management

- Flexible approval model
- Policy enforcement
- Automated lifecycle event management (as user join the company, switch roles, or leave)
- Batch updates (allow an IT team to make changes fo a set of users)
- Automated provisioning to connected applications

#### Compliance Management

- Functionality for applying security standards
- User's access control
- Policy scanning and violation detection

## Identities: Core of Security

> **Identities** represent the users who have access to your corporate system and data. They include anyone or anything that has access to your enterprise systems, including employees, contractors, partners, even robotic users.

IdentityIQ represents each user with an **Identity cube**. It stores all the data collected and used by IdentityIQ for a single user, including identity attributes, enterprise accounts, and type of access held by the user.

### Attribute

A property of an identity, account, application, or entitlment. Attributes often drive IdentityIQ processes. For example, a user's location may be used to give them application access.

### Entitlement

Also known as **permissions**, are the type of access a user has when logging into an application.

### Aggregation

The process of collecting or reading data from applications in your organization. Regular, periodic aggregation is key to keeping identity cubes up to date.

#### Governance Model Steps

1. **Aggregation**: Process of collecting data from the applications in your origanization, the key to keeping identities up-to-date
2. **Compliance Management**: With this module, the enterprise can confirm that the users have onlty the access they need
3. **Lifecycle Management**: Changes management to identities over time

## Access Request

Users can ask for changes to access by requesting the addition or removal of roles and entitlements. Access requests are a function of the IdentityIQ Lifecycle Manager module.

## Event-Driven Provisioning

> **Event-driven provisioning** is an automated process used to keep an identity's access current with their job needs and responsibilities. With it, enterprises can automatically provision or change access that’s required and remove access that's no longer necessary.

### Event-Driven Provisioning Process

- IdentityIQ monitors data: looking for changes, also known as **events**
- IdentityIQ detects changes: when an event, or change, is detected by IdentityIQ, provisioning is initiated
- IdentityIQ provisioning data: the changes are written to the systems or applications that are defined to IdenitityIQ

## Certification and Access Reviews

> **Certification**, also called access certification, is the process of collecting and reviewing the access that identities hold on IT-controlled systems in your organization.

During the certification process, a user's set of access is reviewed by certifier, often a manager or application or entitlement owner. The certifier makes the decision to approve or reject the access that a user currently holds. Certifications should be performed regularly, such as quarterly or annually, to ensure users have only the access they need.

### Certification Steps

Steps of a sample manager certification campaign:

#### 1) A certification campaign is created

A compliance officer creates a manager certification campaign to require all managers to review and certify the access of their direct reports. The campaign creator also specifies when the campaign will run, how long it will last, and other key dates.

#### 2) Access information is collected

IdentityIQ collects and compiles all the required information related to identities, points of access, applications, and more. This process can take minutes, hours, or even days depending on the size of the campaign.

#### 3) Access reviews are created for certifier review

During the generation phase, an access review is created for each certifier designated in the certification campaign, and you can configure IdentityIQ to notify the certifiers that they have access reviews awaiting their action. A typical certification campaign may require certifiers to complete their reviews in a 2 or 4 week timeframe.

#### 4) Each manager approves or revokes access

Managers open their assigned access reviews and approve or revoke access for their employees. If AI-Driven Access Recommendations is deployed in the organization, managers will see a suggestion, based on comparisons with other identities, for each access.

#### 5) Challenge, sign-off, and revocation

An optional challenge phase allows users who would lose access to challenge that decision, and certifiers can reconsider their decisions.

Once finalized, the managers sign-off on the access review, then IdentityIQ begins the revocation process, as needed.

## Policies

> An IdentityIQ policy defines user access conditions that are unwanted by the organization. IdentityIQ policies define the access business policies of your enterprise. First, you define a policy, then IdentityIQ can prevent that condition from occurring or check the identities for that condition.

### Separation of duties (SoD)

Sod policies ensure that users do not hold conflicting access; these policies are very common in many organizations. Generally, SoD policies ensure that no single user has access or responsibility for entire process, such as approving new vendors and also paying those vendors.

### Dormant accounts

Accounts that are unused for excessive lengths of time, can lead to two problems:

- They can cost an organization money for application accounts that aren't being used and they can become a security risk
- If no one is using an account, no one will notice if it is compromised

### Acting on Violations

The person assigned as the violation owner can take action on the violation through an assigned work item. Depending upon the type of violation, there are three possible actions that the owner can take:

- Allow it
- Correct it
- Certify the identity. Administrators have this option, which allows them to take a look at all access held by the user and make decisions about that access holistically

A key purpose of the policy detection process is to ensure that someone is aware of undesired access and can take action to fix the problem. Once you've seen the violation, you can take steps to address it, which may involve actions outside of IdentityIQ. In IdentityIQ, you can allow the violation as an exception for a certain period of time. If the problem still exists after that time period ends, IdentityIQ will flag it again so it isn't forgotten.
