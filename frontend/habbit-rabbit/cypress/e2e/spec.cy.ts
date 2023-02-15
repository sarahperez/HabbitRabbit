describe('Calendar Page Test', () => {
  it('Navigates to the Calendar Page', () => {
    cy.visit('http://localhost:4200/')

    cy.get('.side-nav-item').eq(1).click();
    cy.location('pathname').should('match', /\/calendar$/);
    //cy.contains('full-calendar').should('be.visible');
  })
})