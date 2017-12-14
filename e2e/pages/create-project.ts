import { browser, by, element, promise, ElementFinder } from 'protractor';
import * as utils from '../utils/utils';

export interface Project {
  Name: string;
  Author: string; // for testing
  Mode: string; // go, node etc
  Infra: string[]; // mysql etc
}

export class CreateProject {
  create(p: Project) {
    utils.e('#mode-picker-' + p.Mode).click();
    p.Infra.forEach((v: string) => {
      utils.e('#infra-picker-' + v).click();
    });
    utils.type('#create-project-name', p.Name);
    utils.e('#create-project-submit').click();
  }
}
