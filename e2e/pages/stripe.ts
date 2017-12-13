import { browser, by, element, promise, protractor } from 'protractor';
import * as utils from '../utils/utils';

export class StripePage {
  navigateTo() {
    return browser.get('/none');
  }

  // expects the payment popup in progress of upcoming
  pay(): promise.Promise<void> {
    const until = protractor.ExpectedConditions;
    const stripeSubmitButton = element(by.css('button.Button'));
    return browser.sleep(5000).then(() => {
      browser
        .actions()
        .sendKeys('awewewe87w878a7sd@asasax.com')
        .perform();
      browser
        .actions()
        .sendKeys(protractor.Key.TAB + '4242424242424242')
        .perform();
      browser
        .actions()
        .sendKeys(protractor.Key.TAB + '1122')
        .perform();
      browser
        .actions()
        .sendKeys(protractor.Key.TAB + '111' + protractor.Key.ENTER)
        .perform();

      // this does not seem to work with the stripe popup...
      //  emailInput = utils.e('form.Modal-form input[placeholder="Email"]');
      //  numberInput = utils.e(
      // rm.Modal-form input[placeholder="Card number"]'
      //
      //  expiryInput = utils.e(
      // rm.Modal-form input[placeholder="MM / YY"]'
      //
      //  cvc = utils.e('form.Modal-form input[placeholder="CVC"]');
      // .type(emailInput, 'a@a.com');
      // .type(numberInput, '4242424242424242');
      // .type(expiryInput, '1122');
      // .type(cvc, '111');
      // .e('form.Modal-form button[type=submit]').click();
    });
  }
}
