describe('Calendar Page Test', () => {
  it('Registers a user', () => {
    cy.visit('http://localhost:4200/')
    /*
    cy.get('.username-input').type('sarahperez14');
    cy.get('.password-input').type('C0tt0nc4ndy!');

    cy.get('.sign-in-button').click();
    */
    cy.get('.register-button').click();
    cy.location('pathname').should('match', /\/register$/);
    //cy.contains('full-calendar').should('be.visible');

    cy.get('.name-input').type('Sarah');
    cy.get('.email-input').type('sarahemail@gmail.com');
    cy.get('.username-input').type('SarahUsername');
    cy.get('.password-input').type('sarahpassword123!');

    cy.get('.option-button').click();

    /*
    cy.get('.username-input').type('SarahUsername');
    cy.get('.password-input').type('ssarahpassword123!');

    cy.get('.sign-in-button').click().first();
    */
  })
})