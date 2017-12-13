import { browser, by, element, promise, protractor } from 'protractor';
import { LoginPage, UserReg, UserLogin } from '../pages/login.po';
import * as utils from '../utils/utils';

export class Header {
  register(user: UserReg): promise.Promise<void> {
    expect(this.hasRegisterButton()).toBeTruthy();
    this.clickRegister();
    const loginPage = new LoginPage();
    const until = protractor.ExpectedConditions;
    const regSubmit = utils.e('#register-submit');
    return browser
      .wait(
        until.presenceOf(regSubmit),
        5000,
        'Register tab taking too long to appear in the DOM'
      )
      .then(() => {
        expect(regSubmit.isPresent()).toBeTruthy();
        loginPage.register(user);
        return browser.wait(
          until.titleContains(user.UserName),
          5000,
          'Redirect to author page after register took too long'
        );
      });
  }

  login(user: UserReg): promise.Promise<void> {
    expect(this.hasLoginButton()).toBeTruthy();
    const until = protractor.ExpectedConditions;
    const loginSubmit = utils.e('#login-submit');
    const loginPage = new LoginPage();

    this.clickLogin();
    return browser.wait(until.presenceOf(loginSubmit), 5000).then(() => {
      expect(loginSubmit).toBeTruthy();

      loginPage.login({
        Email: user.Email,
        Password: user.Password
      });
      return browser.wait(
        until.titleContains(user.UserName),
        3000,
        'Redirect to author page after login took too long'
      );
    });
  }

  logout(): promise.Promise<void> {
    const until = protractor.ExpectedConditions;
    const menu = element(by.id('dropdown-menu'));
    const logout = element(by.css('.logout'));
    expect(menu.isPresent()).toBeTruthy();

    this.clickMenu();
    return browser.wait(until.elementToBeClickable(logout), 3000).then(() => {
      expect(logout.isPresent()).toBeTruthy();
      return this.clickLogout().then(() => {
        return browser.wait(until.titleContains('1Backend'), 10000);
      });
    });
  }

  private hasRegisterButton(): promise.Promise<boolean> {
    const regButton = element(by.id('register-button'));
    return regButton.isPresent();
  }

  private clickRegister(): promise.Promise<void> {
    return element(by.id('register-button')).click();
  }

  private hasLoginButton(): promise.Promise<boolean> {
    const loginButton = element(by.id('login-button'));
    return loginButton.isPresent();
  }

  private clickLogin(): promise.Promise<void> {
    return element(by.id('login-button')).click();
  }

  private hasMenu(): promise.Promise<boolean> {
    const menuButton = element(by.id('dropdown-menu'));
    return menuButton.isPresent();
  }

  private clickMenu(): promise.Promise<void> {
    return element(by.id('dropdown-menu')).click();
  }

  private hasLogout(): promise.Promise<boolean> {
    const logoutButton = element(by.css('.logout'));
    return logoutButton.isPresent();
  }

  private clickLogout(): promise.Promise<void> {
    return element(by.css('.logout')).click();
  }
}
