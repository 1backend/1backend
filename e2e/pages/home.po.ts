import { browser, by, element, promise } from 'protractor';

export class HomePage {
  navigateTo() {
    return browser.get('/');
  }

  hasRegisterButton(): promise.Promise<boolean> {
    const regButton = element(by.id('register-button'));
    return regButton.isPresent();
  }

  clickRegister(): promise.Promise<void> {
    return element(by.id('register-button')).click();
  }
}
