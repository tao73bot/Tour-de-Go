
**Evaluation Question: Building a Micro SaaS Backend in Go**

**Objective:**
Design and implement a backend system for a Micro SaaS platform using Go (Golang). The goal of this exercise is to evaluate your skills in backend development, API design, user authentication, role-based access control, and email integration. Please answer the questions and implement the requested features as specified below.

---

### **Scenario:**
You are tasked with building the backend for a simple Micro SaaS platform, specifically a lightweight CRM (Customer Relationship Management) system for small businesses. This platform will provide features for managing customer data and tracking interactions. The system will have two user roles:

1. **🔑 User:** Can manage their own tasks (CRUD operations).
2. **🔒 Admin:** Can manage all users and their tasks, including viewing reports and analytics.

The backend should include the following functionalities:

### **📈 CRM-Specific Features**
This lightweight CRM system should include the following capabilities:

1. **🔍 Lead Management**
   - 🔗 Users can create a lead with fields such as `name`, `email`, `phone`, `source`, and `status`.
   - ➡️ Leads can have their status updated (e.g., from `New` to `Qualified` to `Customer`).

2. **📞 Customer Management**
   - 🔍 Users can convert a lead to a customer.
   - 📋 Customers should have additional fields such as `address`, `companyName`, and `notes`.

3. **📋 Interaction Tracking**
   - ✉️ Users can log interactions (e.g., calls, emails, or meetings) with leads or customers.
   - 🕒 Each interaction should include fields like `type`, `date`, `notes`, and `relatedEntityId` (to associate with a lead or customer).

4. **🔎 Search and Filtering**
   - 🔢 Users should be able to search and filter leads and customers based on various fields (e.g., status, source, or name).

---

### **1. User Registration and Authentication**

1.1 🔑 Implement an endpoint for user registration using email and password.
  - 🔒 Ensure strong password validation.
  - ⚡ Store passwords securely (hint: use hashing with bcrypt or similar).

1.2 🔒 Implement an email verification system:
  - ✉️ Send a verification email with a unique token.
  - 🔗 Create an endpoint to verify the token.

1.3 🔐 Build an endpoint for login:
  - 🌍 Use JWT for session management.
  - ⏳ Include expiration and refresh token logic.

---

### **2. Role-Based Access Control**

2.1 🔒 Define separate routes for users and admins:
  - 🌍 Public routes: Registration, login, and email verification.
  - 🔐 Private routes: Task management (users) and user management (admins).

2.2 ⚠️ Implement middleware to verify JWT tokens and enforce role-based access control.

---

### **3. Task and Activity Management**

3.1 ✅ Implement CRUD operations for activities related to leads or customers:
  - 🔄 An activity should have fields: `title`, `description`, `status` (e.g., pending, in progress, completed), `dueDate`, and `associatedEntityId` (to link it to a lead or customer).
  - 🔑 Users should only manage their own activities.

3.2 ⚙️ Admin should have the ability to:
  - 🔗 View all activities associated with any lead or customer.
  - ❌ Edit or delete activities for any user.

---

### **4. Database Design**

4.1 🔍 Design a database schema that efficiently supports the CRM's features:
  - 🔄 Ensure the schema can handle user authentication, lead and customer management, and activity tracking.
  - 🔒 Consider scalability, data consistency, and performance when structuring your tables or collections.

4.2 ⚡ Justify your choice of database type (SQL vs NoSQL):
  - 🌐 Explain why your selected database (e.g., PostgreSQL, MongoDB) fits the requirements of the CRM.
  - 🔑 Highlight the advantages and any trade-offs in your decision.

---

### **5. API Documentation**

5.1 🖊️ Use Swagger (or a similar tool) to document the API:
  - ➡️ Define all endpoints with request and response examples.
  - 🔗 Clearly specify public and private routes.

5.2 ⚠️ Include a detailed error handling strategy:
  - ❓ Provide common error codes and their corresponding messages.

---

### **6. Implementation Guidelines**

6.1 🚀 Use the **Gin** web framework (or any other Go framework of your choice) to build the application.

6.2 ⚖️ Focus on proper error handling instead of testing and deployment:
  - ⚠️ Ensure clear and consistent error messages.
  - 🔒 Avoid exposing sensitive information in error responses.

6.3 🔧 Ensure clean and modular code structure:
  - 🔢 Use packages to separate concerns (e.g., routes, models, services).
  - 🔗 Include comments to explain key decisions.

---

### **Submission Guidelines:**

- 🔗 Provide the complete Go source code in a GitHub repository.
- 🔢 Include a `README.md` file with instructions to set up and run the project locally.
- 📋 Add Swagger documentation as part of the project.

---

### **Evaluation Criteria:**

- **🔍 Code Quality:** Readability, maintainability, and adherence to best practices.
- **⚠️ Error Handling:** Clarity, consistency, and security of error messages.
- **✅ Completeness:** Implementation of all required features.
- **🔄 Design:** Thoughtful API design and database structure.
- **🔒 Documentation:** Quality and usability of Swagger documentation.

Good luck!
