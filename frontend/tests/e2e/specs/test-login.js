module.exports = {
  'test if user can login': browser => {
    browser.init()
      .openHomepage()
      .waitForElementVisible('Header', 10000)
      .click('a[href="/login"]')
      .waitForElementVisible('form[id="login-form"]', 10000)
      .setValue('input[id="email"]', "richardkerkvliet@hotmail.com")
      .setValue('input[id="password"]', "admin")
      // .setValue('input[id="email"]', "test@email.nl")
      // .setValue('input[id="password"]', "s3m1eybkArcvEIjdZyWI")
      .click('button[type="submit"]')
      .waitForElementVisible('div[class="bandsview"]', 10000)
      .end()
  },

  'Should show error of incorrect credentials': browser => {
    browser.init()
      .openHomepage()
      .waitForElementVisible('Header', 10000)
      .click('a[href="/login"]')
      .waitForElementVisible('form[id="login-form"]', 10000)
      .setValue('input[id="email"]', "nietecht@fout.gft")
      .setValue('input[id="password"]', "bestaatniet!")
      .click('button[type="submit"]')
      .assert.visible('h2')
  }
}