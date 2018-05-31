import * as bigInt from 'big-integer';

import { Address } from '@models';

export class Contact {
  address: Address;
  guid: string;
  isNAT: boolean;

  constructor(config: { address: Address; guid?: string; isNAT?: boolean }) {
    const { address, guid, isNAT } = config;

    this.address = address;
    this.guid = guid || Contact.generateGUID();
    this.isNAT = isNAT || false;
  }

  static generateGUID() {
    return bigInt.randBetween(0, 2 ** 64).toString();
  }
}
