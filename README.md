# OLA 3 - Trailer Rental System

### Our approach to DDD

Event Storming

Key Events to Consider:

- `TrailerRented:` When a customer rents a trailer.
- `TrailerReturned:` When a customer returns a trailer.
- `LateFeeApplied:` When a trailer is returned late.
- `InsurancePurchased:` When a customer buys insurance.
- `PaymentProcessed:` When a rental fee or insurance payment is completed.
- `PartnershipFeeCharged:` When the partner company is charged for the service.

Commands:

- `BookTrailer`
- `ReturnTrailer`
- `ApplyLateFee`
- `ProcessPayment`
- `PurchaseInsurance`
- `ChargePartnershipFee`

Aggregates:

- `Rental:` Represents the rental transaction.
- `Trailer:` Represents a trailer available for rent.
- `Customer:` Represents the user renting a trailer.
- `PartnerCompany:` Represents the business partners who offer locations for the trailers.

### Strategic and Tactical Design Approaches

**Strategic Design**
Bounded Contexts:

- Rental Context: Manages trailer bookings, returns, and rental agreements.
- Payment Context: Manages payments, insurance purchases, and partner company billing.
- Partner Context: Manages partnerships and the business relationships with stores (Jem og Fix, Fog, etc.).
