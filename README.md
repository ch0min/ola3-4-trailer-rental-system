## Trailer Rental System
**Group D**

---

### Repository Structure

This repository contains two branches:
- **`main` branch**: Contains the monolithic solution for the MyTrailer application.
- **`microservice` branch**: Contains the refactored microservice architecture solution, where the monolithic application has been split into services.

### How to Navigate the Project

1. **Monolith Solution (Main Branch)**:
   - The `main` branch contains the initial implementation of the Trailer Rental System as a **monolithic application**. All components (rental management, user management, payments, etc.) are part of a single application codebase.
   - **How to run**: Follow the instructions in the `README` on the `main` branch for setup and running the monolith.

2. **Microservice Solution (Microservice Branch)**:
   - The `microservice` branch demonstrates how we refactored the monolithic solution into separate **microservices**. Each service (rental, user, payment, etc.) runs independently and communicates with other services through APIs and event-driven messaging.
   - **How to run**: Follow the `README` on the `microservice` branch to understand how to set up and run the individual services using Docker or Docker Compose.

---

### Technology Stack

**Version Control Platform**:  
- Git - GitHub

**Text Editing and Development Environment**:  
- VScode/Neovim  
- DBeaver  
- Postman

**General Online Research Tools**:  
- Stack Overflow  
- MDN Web Docs  
- Golang Docs

### Development Stack

**Backend Development**:  
- Golang

**Frontend Development**:  
- TypeScript - (React Framework)

**Database Management**:  
- PostgreSQL

**Development Tools**:  
- Docker  
- Docker Compose

**CI/CD Pipeline**:  
- GitHub Actions

---

### Strategic Design

**Core Domain**:  
The core of MyTrailer is the trailer rental process, including booking, insurance, and payments.

**Subdomains**:
- **Rental Management**: Handles trailer reservations, bookings, returns, and excess fees for late returns.
- **Payment and Insurance**: Manages payments for insurance and partnerships with companies for trailer placements.
- **Location Management**: Tracks trailer locations across different companies (e.g., Jem og Fix, Fog).
- **Company Partnerships**: Manages agreements with the companies offering trailer parking space.

### Bounded Context Diagram
This diagram shows the different bounded contexts and their relationships.

![image](docs/models/bcd.png)

---

### Tactical Design with DDD

**Entities, Value Objects, and Aggregates**:

- **Aggregate: Rental**
  - Properties: `startTime`, `endTime`, `rentalFee`, `excessFee`

- **Entity: Trailer**
  - Properties: `number`, `availabilityStatus`

- **Entity: Customer**
  - Properties: `name`

- **Value Object: ContactDetails**
  - Properties: `email`, `phoneNumber`, `address`

- **Value Object: Location**
  - Properties: `locationName`, `locationAddress`

![image](docs/models/domain.png)

---

### Event-Driven Architecture (Messaging Flow)

Since the system has several independent components (rental, payment, partnerships), using an event-driven architecture could help:

**Events**:
- `TrailerBooked`: Triggers when a user books a trailer.
- `TrailerReturned`: Triggers when a trailer is returned. It updates availability and calculates any excess fees.
- `PaymentProcessed`: Triggers after rental completion, charging the customer (for insurance or excess fees).
- `PartnershipPaymentTriggered`: Occurs when a partnership agreement triggers a payment between MyTrailer and store partners.

---

### C4 Diagrams

We chose the **Context** and **Component** diagrams because they provide the right balance of abstraction for this project. The Context diagram helps us visualize how the system interacts with external actors, while the Component diagram shows the internal structure of key services. This allows us to focus on system design without getting too detailed or technical, making it suitable for this small project within the given timeframe.

- **System Architecture**: The MyTrailer system consists of several services, including rental, user and with further development payment, and partnership services.
![image](docs/models/architecture_diagram.png)

- **Context Diagram**: Illustrates how the MyTrailer system interacts with external entities like the app, payment gateways, and store partners.
![image](docs/models/c4_cntx.png)

---

- **Component Diagram**: Shows the internal services of the system, such as `RentalService`, `PaymentService`, `PartnershipService`, etc.
![image](docs/models/c4_comp.png)

---

### Behavior-Driven Development (BDD)
Example of BDD usage in this project:

**Scenario: Customer books a trailer**  
- **Given** a customer is logged into the MyTrailer app  
- **When** they select a location and book a trailer  
- **Then** the system confirms the booking  
- **And** updates the availability of the trailer

---

### Ubiquitous Language

- **Trailer**: The rental item placed at partner locations.
- **Rental**: A short-term or long-term booking of a trailer.
- **Excess Fee**: The penalty for late returns.
- **Insurance**: An optional fee for protection against damage.
- **Location**: The specific place where a trailer is parked.

---

### Definition of DONE

- **Model Completeness**: When the core entities (`Trailer`, `Rental`, `Customer`, `Location`) and their behaviors are fully modeled.
- **Messaging Flow**: When event-driven messages (`TrailerBooked`, `PaymentProcessed`, etc.) are identified and mapped between services.
- **Architecture Documentation**: When all the diagrams (C4, component, sequence) are completed.
- **Testing Ready**: The system is ready to start experimenting with BDD scenarios and the flow of events between services.

---

### Instructions to Run

1. **Monolith Solution**: (On `main` branch)
   - Clone the repository.
   - Follow the `README.md` on the `main` branch for running the monolithic solution.

2. **Microservice Solution**: (On `microservice` branch)
   - Switch to the `microservice` branch: `git checkout microservice`.
   - Follow the `README.md` on the `microservice` branch to run the individual services.

---
