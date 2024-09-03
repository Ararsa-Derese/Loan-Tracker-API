## Functional Requirements

### 1. User Management

#### 1.1 User Registration
- **Description:** Users can register on the platform with their email, password, and profile details.
- **Actors:** Visitor
- **Flow of Events:**
  1. The user submits necessary details (username, email, password, etc.).
  2. The system validates the information (email format, password strength, etc.).
  3. The system checks for existing users with the same email or username.
  4. If validation passes, a new user account is created.
  5. The system may send an activation email or link for account verification.

#### 1.2 Login
- **Description:** Registered users can log in to access the platform's features.
- **Actors:** User & Admin
- **Flow of Events:**
  1. The user provides their username (or email) and password.
  2. The system validates the credentials.
  3. If valid, the system generates and returns access and refresh tokens.

#### 1.3 Authentication
- **Description:** Allows users to access the system without needing to log in every time.
- **Actors:** User & Admin
- **Flow of Events:**
  1. The user sends the access token with each request.
  2. The server verifies the token.
  3. If expired, the system validates the refresh token and issues a new access token.

#### 1.4 Forgot Password
- **Description:** Users can reset their password if forgotten.
- **Actors:** User & Admin
- **Flow of Events:**
  1. User requests a password reset by providing their email.
  2. The system sends a password reset link via email.
  3. The user resets the password using the link.
  4. The system updates the password and confirms the change.

### 2. User Manipulation
#### 2.1 Get all Users
- **Description:** Admins and root user can get users in the database
- **Actors:** Admin and Root User
- **Flow of Events:**
  1. Admin Logs in using email and password which is secret from other users (you can speficy in the enviroment variables).
  2. Now they can see the list of the users.

  #### 2.2 Delete User
- **Description:** Admins and root user can get users in the database.
- **Actors:** Admin & Root user
- **Flow of Events:**
  1. Admin Logs in using email and password which is secret from other users (you can speficy in the enviroment variables).
  2. Now they can Delete any user by providing userID of the users.


