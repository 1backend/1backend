import { browser, by, element, promise } from 'protractor';

export class Header {
  hasRegisterButton(): promise.Promise<boolean> {
    const regButton = element(by.id('register-button'));
    return regButton.isPresent();
  }

  clickRegister(): promise.Promise<void> {
    return element(by.id('register-button')).click();
  }
  hasLoginButton(): promise.Promise<boolean> {
    const loginButton = element(by.id('login-button'));
    return loginButton.isPresent();
  }

  clickLogin(): promise.Promise<void> {
    return element(by.id('login-button')).click();
  }
  hasMenu(): promise.Promise<boolean> {
    const menuButton = element(by.id('dropdown-menu'));
    return menuButton.isPresent();
  }
  clickMenu(): promise.Promise<void> {
    return element(by.id('dropdown-menu')).click();
  }
  hasLogout(): promise.Promise<boolean> {
    const logoutButton = element(by.css('.logout'));
    return logoutButton.isPresent();
  }
  clickLogout(): promise.Promise<void> {
    return element(by.css('.logout')).click();
  }
}
