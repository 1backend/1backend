import { HomePage } from './pages/home.po';
import { LoginPage } from './pages/login.po';
import {
  browser,
  by,
  element,
  promise,
  ElementFinder,
  protractor
} from 'protractor';

function makeid() {
  let text = '';
  const possible = 'abcdefghijklmnopqrstuvwxyz0123456789';
  for (let i = 0; i < 5; i++) {
    text += possible.charAt(Math.floor(Math.random() * possible.length));
  }
  return text;
}

describe('Home page', () => {
  let homePage: HomePage;
  let loginPage: LoginPage;
  const id = makeid();

  browser.driver
    .manage()
    .window()
    .setSize(1200, 800);

  beforeEach(() => {
    homePage = new HomePage();
    loginPage = new LoginPage();
  });

  it('should be able to register', () => {
    browser.waitForAngularEnabled(false);
    homePage.navigateTo();

    expect(homePage.hasRegisterButton()).toBeTruthy();
    homePage.clickRegister();

    const until = protractor.ExpectedConditions;
    const regSubmit = element(by.id('register-submit'));
    browser
      .wait(
        until.presenceOf(regSubmit),
        3000,
        'Register tab taking too long to appear in the DOM'
      )
      .then(() => {
        expect(regSubmit.isPresent()).toBeTruthy();
        loginPage.register({
          UserName: 'user-' + id,
          Email: 'user' + id + '@gmail.com',
          Password: 'pw'
        });
        return browser.wait(
          until.titleContains('user-' + id),
          3000,
          'Redirect to author page took too long'
        );
      });
  });
});
