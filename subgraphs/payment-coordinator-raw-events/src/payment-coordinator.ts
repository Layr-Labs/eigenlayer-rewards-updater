import { RangePaymentCreated as RangePaymentCreatedEvent } from "../generated/PaymentCoordinator/PaymentCoordinator"
import { RangePaymentCreated } from "../generated/schema"

export function handleRangePaymentCreated(
  event: RangePaymentCreatedEvent
): void {
  let entity = new RangePaymentCreated(
    event.transaction.hash.concatI32(event.logIndex.toI32())
  )
  entity.rangePayment_avs = event.params.rangePayment.avs
  entity.rangePayment_strategy = event.params.rangePayment.strategy
  entity.rangePayment_token = event.params.rangePayment.token
  entity.rangePayment_amount = event.params.rangePayment.amount
  entity.rangePayment_startRangeTimestamp =
    event.params.rangePayment.startRangeTimestamp
  entity.rangePayment_endRangeTimestamp =
    event.params.rangePayment.endRangeTimestamp

  entity.blockNumber = event.block.number
  entity.blockTimestamp = event.block.timestamp
  entity.transactionHash = event.transaction.hash

  entity.save()
}
