import * as bigInt from 'big-integer';
import {
  always,
  equals,
  find,
  flatten,
  pipe,
  propEq,
  reject,
  slice,
  sort
} from 'ramda';

import { Contact } from '@models/contact';
import logger from '@utils/logging';

export class RoutingTable {
  private buckets: Contact[][];

  static readonly bucketCount = 64;
  static readonly bucketSize = 10;

  constructor(private selfNode: Contact) {
    this.buckets = Array2D<Contact>(RoutingTable.bucketCount);
  }

  addNode(node: Contact): void {
    const bucket = this.selectBucket(node.guid);

    const notSelf = node.guid !== this.selfNode.guid;
    const bucketNotFull = this.buckets[bucket].length < RoutingTable.bucketSize;

    if (notSelf && bucketNotFull) {
      logger.info(`P2P layer: Adding node ${node.guid} to routing table.`);
      this.buckets[bucket] = [...this.buckets[bucket], node];
      logger.info(
        `RoutingTable: ${JSON.stringify(
          this.buckets.filter((x) => x.length),
          null,
          2
        )}`
      );
    }
  }

  removeNode(node: Contact): void {
    logger.info(`P2P layer: Removing node ${node.guid} from routing table.`);
    let bucket = this.selectBucket(node.guid);
    this.buckets[bucket] = reject(equals(node), this.buckets[bucket]);
  }

  getAllNodes(): Contact[] {
    return flatten(this.buckets);
  }

  getNodeByGUID(guid: string): Contact | undefined {
    return find(propEq('guid', guid), this.buckets[this.selectBucket(guid)]);
  }

  getNearestNodes(guid: string, limit = RoutingTable.bucketSize): Contact[] {
    const nearestNodes = pipe<
      Contact[][],
      Contact[],
      Contact[],
      Contact[],
      Contact[]
    >(
      flatten,
      reject(propEq('guid', guid)),
      sort((node) =>
        bigInt(node.guid)
          .xor(bigInt(guid))
          .compare(bigInt(guid))
      ),
      slice(0, limit)
    )(this.buckets);

    return nearestNodes;
  }

  private selectBucket(guid: string): number {
    const xor = bigInt(guid).xor(bigInt(this.selfNode.guid));
    return xor.bitLength().toJSNumber() - 1;
  }
}

function Array2D<T>(size: number): T[][] {
  return Array(size)
    .fill(0)
    .map(always([]));
}
