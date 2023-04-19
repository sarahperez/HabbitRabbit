/*
describe('register', () => {
  it('Registers a user', () => {
    cy.visit('http://localhost:4200/')
  
    cy.get('.register-button').click();
    cy.location('pathname').should('match', /\/register$/);
    //cy.contains('full-calendar').should('be.visible');

    cy.get('.name-input').type('Sarah');
    cy.get('.email-input').type('sarahemail@gmail.com');
    cy.get('.username-input').type('SarahUsername');
    cy.get('.password-input').type('sarahpassword123!');

    cy.get('.option-button').click();

    //register service should bring the user back to the login page
    cy.location('pathname').should('match', /\/login$/);

    
    //Newly registered user should be able to login now
    cy.get('.username-input').type('SarahUsername');
    cy.get('.password-input').type('sarahpassword123!');

    cy.get('.sign-in-button').click().first();

    //User is brought to the home page
    cy.location('pathname').should('match', /\/home$/);
    
  })
})
*/
describe('Login page', () => {
  it('Logs the user in', () => {
    cy.visit('http://localhost:4200/');
    
    cy.get('.username-input').type('sarahperez14');
    cy.get('.password-input').type('C0tt0nc4ndy!');
    cy.get('.sign-in-button').click();
    cy.location('pathname').should('match', /\/login$/);
  })
})

describe('Calendar Page', () => {
  it('Brings the user to calendar page', () => {
    cy.visit('http://localhost:4200/home');
    cy.get('.side-nav-nav-link').click({multiple: true});
    cy.visit('http://localhost:4200/calendar');
    cy.location('pathname').should('match', /\/calendar$/);
  })

  it('Adds an event to the calendar', () => {
    cy.visit('http://localhost:4200/calendar');
  })

  it('Deletes an event from the calendar', () => {

  })
})

describe('Todo Page', () => {
  it('Brings the user to todo page', () => {
    cy.visit('http://localhost:4200/home');
    cy.get('.side-nav-nav-link').click({multiple: true});
    cy.visit('http://localhost:4200/calendar');
    cy.location('pathname').should('match', /\/calendar$/);
  })

  it('Adds  task to the todo list', () => {
    cy.visit('http://localhost:4200/calendar');
  })

  it('Deletes a task from the todo list', () => {

  })
})