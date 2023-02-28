describe('Calendar Page Test', () => {
  it('Navigates to the Calendar Page', () => {
    cy.visit('http://localhost:4200/')

    cy.get('.username-input').type('sarahperez14');
    cy.get('.password-input').type('C0tt0nc4ndy!');

    cy.contains('Sign').click();
    cy.location('pathname').should('match', /\/home$/);
    //cy.contains('full-calendar').should('be.visible');
  })
})