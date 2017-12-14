import { browser, by, element, promise, protractor } from 'protractor';
import * as utils from '../utils/utils';
import { StripePage } from '../pages/stripe';
import { CreateProject, Project } from '../pages/create-project';

export class AuthorPage {
  pricePer100k = 9;
  name = '';

  constructor(authorName: string) {
    this.name = authorName;
  }

  navigateTo() {
    return browser.get('/' + this.name); // this does not work
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

  createProject(p: Project): promise.Promise<any> {
    const until = protractor.ExpectedConditions;
    const addProjectButton = utils.e('#add-project-button');
    return browser
      .wait(
        until.presenceOf(addProjectButton),
        9000,
        'Loading author page took too long'
      )
      .then(() => {
        expect(addProjectButton.isPresent()).toBeTruthy();
        addProjectButton.click();

        const createProjectButton = utils.e('#create-project-submit');
        return browser.wait(
          until.presenceOf(createProjectButton),
          9000,
          'Create project popup taking too long to come up'
        );
      })
      .then(() => {
        const cp = new CreateProject();
        cp.create(p);
        return browser.wait(
          until.invisibilityOf(utils.e('#project-status-red')),
          9000,
          'Create project build took too long'
        );
      })
      .then(() => {
        return browser.wait(
          until.invisibilityOf(utils.e('#build-tab-icon-inprogress')),
          45000,
          'Project build took too long'
        );
      })
      .then(() => {
        return browser.wait(
          until.presenceOf(utils.e('#build-tab-icon-success')),
          1000,
          'Build is unsuccessful'
        );
      });
  }
}
