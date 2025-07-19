# Active Context - Chess Tournament Hub

## Current Work Focus

### Immediate Priorities (Next 2-4 Weeks)
1. **Project Foundation Setup**
   - Finalize technical stack decisions (Go backend confirmed)
   - Set up development environment and repository structure
   - Establish CI/CD pipeline for multiple repositories
   - Create development team structure and workflows

2. **Architecture Finalization**
   - Complete database schema design for core entities
   - Finalize API contract specifications
   - Set up authentication and authorization framework
   - Plan payment integration architecture with CHIP Asia (https://www.chip-in.asia)

3. **Sprint 1 Preparation**
   - Break down Phase 1 deliverables into actionable tasks
   - Set up project management tools and task tracking
   - Establish code review processes and quality standards
   - Plan team onboarding and knowledge sharing sessions

### Current Sprint Goals (Sprint 1: Weeks 1-2)
- **Backend Foundation:** Go project setup with Gin, PostgreSQL integration
- **Authentication System:** JWT-based auth with user registration/login
- **Frontend Foundation:** Next.js setup with authentication pages
- **Development Environment:** Docker setup for local development
- **Repository Structure:** 4-repo setup (API, Web, Mobile, Infrastructure)

## Recent Decisions

### Key Technical Decisions Made
1. **Backend Language:** Go with Gin framework (chosen over Node.js for performance)
2. **Database:** PostgreSQL as primary database with Redis for caching
3. **Frontend:** Next.js with React and TypeScript
4. **Mobile:** React Native (planned for Phase 2)
5. **Payments:** CHIP Collect and CHIP Send with escrow system (platform as payment processor)
6. **Architecture:** Multi-repository approach for independent deployment cycles

### Business Model Decisions
1. **Revenue Model:**
   - RM50 and below: RM5 flat fee
   - RM51-RM60: 10% commission fee
   - RM61-RM84: 9% commission fee
   - RM85 and above: 8% commission fee
2. **Market Focus:** Malaysia first, then ASEAN expansion
3. **User Acquisition:** Partnership with local chess clubs and federations
4. **Pricing Strategy:** Commission-based with premium organizer features

### User Experience Decisions
1. **Mobile-First:** Prioritize mobile experience for Southeast Asian market
2. **Registration Flow:** Separate flows for players vs organizers
3. **Payment Security:** Escrow system
4. **Tournament Discovery:** Search and filter-heavy interface

## Next Steps

### Week 1-2 (Sprint 1)
- [X] Set up Go backend project structure with proper module organization
- [ ] Implement PostgreSQL connection and migration system
- [ ] Create user authentication endpoints (register, login, verify email)
- [ ] Set up Next.js frontend with Tailwind CSS
- [ ] Build authentication pages and protected route middleware
- [X] Establish Docker development environment

### Week 3-4 (Sprint 2)
- [ ] Implement user profile management (players vs organizers)
- [ ] Build user dashboard interfaces
- [ ] Add phone verification and profile completion flows
- [ ] Set up basic tournament data models
- [ ] Create tournament creation form (basic version)
- [ ] Implement tournament listing and detail pages

### Week 5-6 (Sprint 3)
- [ ] Complete tournament creation and management APIs
- [ ] Build tournament registration system
- [ ] Implement basic payment integration with CHIP
- [ ] Add email notification system
- [ ] Create organizer dashboard for tournament management

## Active Decisions and Considerations

### Technical Considerations
1. **API Versioning Strategy:** Implementing `/api/v1/` from start for future compatibility
2. **Database Design:** Balancing normalization with query performance for tournament data
3. **Caching Strategy:** Redis for session management and frequently accessed tournament data
4. **Error Handling:** Standardized error responses across all API endpoints
5. **Logging Strategy:** Structured logging with different levels for development vs production

### Payment Architecture Decisions
1. **Escrow Period:** Pending. Need discussion with organizers on how urgent they need the fund.
2. **Refund Policy:** Automated refunds based on cancellation timing rules
3. **Malaysian Compliance:** Early consultation with legal team for financial regulations
4. **Multi-currency Support:** Plan for regional expansion from architecture start

### User Experience Considerations
1. **Registration Complexity:** Balancing comprehensive profiles with signup friction
2. **Tournament Discovery:** Designing filters that match actual player search behavior
3. **Mobile Performance:** Optimizing for slower internet connections in region
4. **Language Localization:** Planning i18n architecture for Phase 3 expansion

## Important Patterns and Preferences

### Code Patterns
1. **Go Project Structure:** Following standard Go project layout with `/cmd`, `/`, `/pkg`
2. **API Design:** RESTful APIs with consistent naming conventions
3. **Error Handling:** Centralized error handling with proper HTTP status codes
4. **Database Queries:** Using prepared statements and proper connection pooling
5. **Testing Strategy:** Unit tests for business logic, integration tests for APIs

### UI/UX Patterns
1. **Component Library:** Building reusable components with consistent design system
2. **Form Validation:** Real-time validation with clear error messaging
3. **Loading States:** Proper loading indicators for all async operations
4. **Mobile Navigation:** Bottom navigation for mobile, sidebar for desktop
5. **Accessibility:** WCAG 2.1 AA compliance from the start

### Development Workflow
1. **Git Strategy:** Feature branches with pull request reviews
2. **CI/CD Pipeline:** Automated testing and deployment for each repository
3. **Code Reviews:** Required reviews for all code changes
4. **Documentation:** API documentation generated from code annotations
5. **Environment Management:** Separate staging and production environments

## Learnings and Project Insights

### Market Research Insights
1. **Chess Community Behavior:** Strong preference for WhatsApp communication in Malaysia
2. **Payment Preferences:** FPX widely used, credit card adoption growing
3. **Mobile Usage:** High mobile usage, especially among younger players
4. **Tournament Frequency:** Most local tournaments are monthly, larger events quarterly

### Technical Insights
1. **Go Performance:** Excellent choice for concurrent tournament registration handling
2. **PostgreSQL Benefits:** JSON columns useful for flexible tournament configuration storage
3. **CHIP Collect & CHIP Send:** Well-suited for marketplace model with Malaysian support
4. **Next.js Advantages:** Server-side rendering beneficial for SEO and tournament discovery

### Business Insights
1. **Network Effects:** Platform value increases exponentially with user base
2. **Trust Building:** Organizer verification and player reviews crucial for adoption
3. **Local Partnerships:** Chess federation relationships key to credibility
4. **Competition Analysis:** Generic event platforms lack chess-specific features

### User Feedback Patterns (from initial research)
1. **Organizers Want:** Reduced administrative work, professional appearance, better reach
2. **Players Want:** Easy discovery, secure payments, tournament history tracking
3. **Common Pain Points:** Last-minute tournament changes, payment disputes, communication gaps
4. **Feature Requests:** Mobile app, rating integration, social features

## Risk Mitigation Strategies

### Technical Risks
1. **Scalability:** Cloud-native architecture with horizontal scaling capabilities
2. **Payment Security:** PCI compliance and fraud detection from day one
3. **Data Protection:** GDPR-style privacy controls and data encryption
4. **API Reliability:** Comprehensive error handling and graceful degradation

### Business Risks
1. **Market Adoption:** Early partnership with key chess organizations
2. **Competition:** Focus on chess-specific features that generic platforms can't match
3. **Regulatory Changes:** Legal consultation and compliance monitoring
4. **Funding:** Bootstrap approach with early revenue focus

### Operational Risks
1. **Team Scaling:** Clear documentation and onboarding processes
2. **Technical Debt:** Regular refactoring and code quality maintenance
3. **Customer Support:** Early investment in support systems and processes
4. **Feature Creep:** Strict adherence to MVP principles and user feedback validation

This active context should be updated weekly to reflect current priorities, recent decisions, and emerging insights as the project progresses through its development phases.
