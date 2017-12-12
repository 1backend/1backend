import { browser, by, element, promise, ElementFinder } from 'protractor';

// shorthand for element by css
export function e(s: string): ElementFinder {
  return element(by.css(s));
}

// type to input field
export function type(id: string, text: string) {
  const el = element(by.id(id));
  el.click();
  browser.sleep(400);
  el.sendKeys(text);
}
