import {
  assert,
  describe,
  test,
  clearStore,
  beforeAll,
  afterAll
} from "matchstick-as/assembly/index"
import { BigInt, Address, Bytes } from "@graphprotocol/graph-ts"
import { ActivationDelaySet } from "../generated/schema"
import { ActivationDelaySet as ActivationDelaySetEvent } from "../generated/ClaimingManager/ClaimingManager"
import { handleActivationDelaySet } from "../src/claiming-manager"
import { createActivationDelaySetEvent } from "./claiming-manager-utils"

// Tests structure (matchstick-as >=0.5.0)
// https://thegraph.com/docs/en/developer/matchstick/#tests-structure-0-5-0

describe("Describe entity assertions", () => {
  beforeAll(() => {
    let oldActivationDelay = BigInt.fromI32(234)
    let newActivationDelay = BigInt.fromI32(234)
    let newActivationDelaySetEvent = createActivationDelaySetEvent(
      oldActivationDelay,
      newActivationDelay
    )
    handleActivationDelaySet(newActivationDelaySetEvent)
  })

  afterAll(() => {
    clearStore()
  })

  // For more test scenarios, see:
  // https://thegraph.com/docs/en/developer/matchstick/#write-a-unit-test

  test("ActivationDelaySet created and stored", () => {
    assert.entityCount("ActivationDelaySet", 1)

    // 0xa16081f360e3847006db660bae1c6d1b2e17ec2a is the default address used in newMockEvent() function
    assert.fieldEquals(
      "ActivationDelaySet",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "oldActivationDelay",
      "234"
    )
    assert.fieldEquals(
      "ActivationDelaySet",
      "0xa16081f360e3847006db660bae1c6d1b2e17ec2a-1",
      "newActivationDelay",
      "234"
    )

    // More assert options:
    // https://thegraph.com/docs/en/developer/matchstick/#asserts
  })
})
