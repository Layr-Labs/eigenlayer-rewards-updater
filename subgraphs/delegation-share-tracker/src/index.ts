import { Bytes, BigInt } from "@graphprotocol/graph-ts"
import {
  OperatorSharesDecreased as OperatorSharesDecreasedEvent,
  OperatorSharesIncreased as OperatorSharesIncreasedEvent,
} from "../generated/DelegationManager/DelegationManager"
import {
  StakerDelegationShare
} from "../generated/schema"

const ZERO_ADDRESS = Bytes.fromHexString('0x0000000000000000000000000000000000000000')

function getOrCreateStakerDelegationShare(staker: Bytes, strategy: Bytes): StakerDelegationShare {
  let id = staker.concat(strategy)
  let stakerDelegationShare = StakerDelegationShare.load(id)
  if (stakerDelegationShare == null) {
    stakerDelegationShare = new StakerDelegationShare(id)
    stakerDelegationShare.staker = staker
    stakerDelegationShare.strategy = strategy
    stakerDelegationShare.shares = BigInt.fromI32(0)
    stakerDelegationShare.updateBlockTimestamp = BigInt.fromI32(0)
  }
  return stakerDelegationShare as StakerDelegationShare
}

export function handleOperatorSharesDecreased(
  event: OperatorSharesDecreasedEvent
): void {
  let stakerDelegationShare = getOrCreateStakerDelegationShare(event.params.staker, event.params.strategy)

  stakerDelegationShare.shares = stakerDelegationShare.shares.minus(event.params.shares)
  if (stakerDelegationShare.shares.equals(BigInt.fromI32(0))) {
    stakerDelegationShare.operator = ZERO_ADDRESS
  }
  stakerDelegationShare.updateBlockTimestamp = event.block.timestamp

  stakerDelegationShare.save()
}

export function handleOperatorSharesIncreased(
  event: OperatorSharesIncreasedEvent
): void {
  let stakerDelegationShare = getOrCreateStakerDelegationShare(event.params.staker, event.params.strategy)

  stakerDelegationShare.shares = stakerDelegationShare.shares.plus(event.params.shares)
  stakerDelegationShare.operator = event.params.operator
  stakerDelegationShare.updateBlockTimestamp = event.block.timestamp

  stakerDelegationShare.save()
}
