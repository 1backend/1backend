import { browser, by, element, promise, protractor } from 'protractor';
import * as utils from '../utils/utils';
import { StripePage } from '../pages/stripe';

export class AuthorPage {
  pricePer100k = 9;

  navigateTo() {
    return browser.get('/author'); // this does not work
  }

  topUpBy(n: number): promise.Promise<any> {
    const topUpButton = element(by.id('top-up'));
    expect(topUpButton.isPresent()).toBeTruthy();
    topUpButton.click();

    const stripePage = new StripePage();

    const callsLeftDiv = utils.e('#all-calls-left');
    let callsLeft: number;
    callsLeftDiv.getAttribute('data-all-calls-left').then(function(value) {
      callsLeft = +value;
    });

    const until = protractor.ExpectedConditions;
    return stripePage.pay().then(() => {
      return browser
        .wait(
          function() {
            return callsLeftDiv
              .getAttribute('data-all-calls-left')
              .then(function(value) {
                return value === callsLeft + 100000 + '';
              });
          },
          9000,
          'Quota does not seem to match'
        )
        .then(() => {
          return browser.sleep(4000);
        });
    });
  }
}
