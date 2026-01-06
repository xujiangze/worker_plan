# Frontend API Response Handling

## ADDED Requirements

### Requirement: Frontend API response interceptor MUST extract data field from backend response

The frontend response interceptor in `frontend/src/utils/request.ts` MUST automatically extract the `data` field from the backend's unified response structure to ensure API calls receive the actual data payload instead of the wrapper object.

#### Scenario: Create plan API returns correct data structure

**Given** the backend returns a response in the format `{code: 0, message: "success", data: {...}}`
**When** the frontend calls `planApi.createPlan(data)`
**Then** the response should be the Plan object from the `data` field, not the entire response wrapper
**And** the store can successfully call `plans.value.unshift(plan)` without errors

#### Scenario: All API calls receive unwrapped data

**Given** the response interceptor extracts `response.data.data`
**When** any API call is made (getPlans, getPlan, updatePlan, deletePlan, etc.)
**Then** the response should be the actual data payload
**And** the store can process the data without manual unwrapping

#### Scenario: Error responses are handled correctly

**Given** the backend returns an error response with a non-zero code
**When** the API call fails
**Then** the error should be caught and handled by the error interceptor
**And** the user should see an appropriate error message

### Requirement: Frontend store MUST handle API responses correctly

The frontend stores MUST correctly handle API responses after the response interceptor modification, ensuring all CRUD operations work as expected.

#### Scenario: Create plan adds new plan to list

**Given** a user fills out the create plan form
**When** the user submits the form
**Then** the new plan should be added to the beginning of the plans list
**And** the total count should be incremented
**And** a success message should be displayed

#### Scenario: Update plan modifies existing plan in list

**Given** a user edits an existing plan
**When** the user saves the changes
**Then** the plan should be updated in the plans list
**And** if the plan is currently displayed, it should be updated
**And** a success message should be displayed

#### Scenario: Delete plan removes plan from list

**Given** a user deletes a plan
**When** the deletion is confirmed
**Then** the plan should be removed from the plans list
**And** the total count should be decremented
**And** if the plan was currently displayed, it should be cleared
**And** a success message should be displayed

#### Scenario: Fetch plans populates list correctly

**Given** the user navigates to the plans list page
**When** the page loads
**Then** the plans list should be populated with data from the API
**And** pagination information should be updated
**And** the loading state should be cleared

### Requirement: Frontend UI MUST display plans correctly

The frontend UI components MUST correctly display plan data after the API response handling fix.

#### Scenario: Plan list displays all plans

**Given** the plans list is loaded
**When** the user views the list
**Then** all plans should be displayed with correct information
**And** pagination controls should work correctly
**And** filters should work correctly

#### Scenario: Plan form displays plan details

**Given** a user opens the plan form in edit mode
**When** the form loads
**Then** the form should be pre-filled with the plan's current data
**And** all fields should display correctly

#### Scenario: No errors in browser console

**Given** the user interacts with the application
**When** any API call is made
**Then** there should be no errors in the browser console
**And** all operations should complete successfully
