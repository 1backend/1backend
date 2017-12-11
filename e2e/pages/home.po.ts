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
    const regButton = element(by.id('register-button'));
    return browser.driver
      .wait(regButton.isPresent(), 2000)
      .then(() => {
        alert('fasz');
        return regButton.click();
      })
      .catch(err => {
        console.log(err);
        alert(err);
      });
  }
}
