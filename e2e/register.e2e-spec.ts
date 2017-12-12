import { HomePage } from './pages/home.po';
import { LoginPage } from './pages/login.po';
import { Header } from './pages/header.po';
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
  const possible = 'abcdfghijklmnopqstuvwxyz012346789';
  for (let i = 0; i < 5; i++) {
    text += possible.charAt(Math.floor(Math.random() * possible.length));
  }
  return text;
}

describe('Home page', () => {
  let homePage: HomePage;
  let loginPage: LoginPage;
  let header: Header;
  const id = makeid();

  browser.driver
    .manage()
    .window()
    .setSize(1200, 800);

  beforeAll(() => {
    homePage = new HomePage();
    loginPage = new LoginPage();
    header = new Header();
    const Id = this.id;
  });

  it('should be able to register', () => {
    browser.waitForAngularEnabled(false);
    homePage.navigateTo();

    expect(header.hasRegisterButton()).toBeTruthy();
    header.clickRegister();

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
  it('should be able to login', () => {
    const until = protractor.ExpectedConditions;
    const menu = element(by.id('dropdown-menu'));
    const logout = element(by.css('.logout'));
    browser.wait(until.presenceOf(menu), 10000).then(() => {
      expect(menu.isPresent()).toBeTruthy();
      header.clickMenu();
      browser.wait(until.elementToBeClickable(logout), 10000).then(() => {
        expect(logout.isPresent()).toBeTruthy();
        header.clickLogout();
      });
      browser.waitForAngularEnabled(false);
      return browser.wait(until.titleContains('1Backend'), 10000);
    });
    expect(header.hasLoginButton()).toBeTruthy();
    header.clickLogin();
    const loginSubmit = element(by.id('login-submit'));
    browser.wait(until.presenceOf(loginSubmit), 5000).then(() => {
      expect(loginSubmit).toBeTruthy();
      loginPage.login({
        Email: 'user' + id + '@gmail.com',
        Password: 'pw',
        UserName: null
      });
      return browser.wait(
        until.titleContains('user-' + id),
        3000,
        'Redirect to author page took too long'
      );
    });
  });
});
