import { browser, by, element, promise, ElementFinder } from 'protractor';

// shorthand for element by css
export function e(s: string): ElementFinder {
  return element(by.css(s));
}

// type to input field
export function type(elem: string | ElementFinder, text: string) {
  let el: ElementFinder;
  if (typeof elem === 'string') {
    el = element(by.css(<string>elem));
  } else {
    el = <ElementFinder>elem;
  }
  el.click();
  browser.sleep(400);
  el.sendKeys(text);
}
