/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "kkaddal.omni.ethereumbridge";

export interface EthState {
  index: string;
  address: string;
  slot: string;
  balance: string;
  storage: string;
  blocknumber: string;
}

const baseEthState: object = {
  index: "",
  address: "",
  slot: "",
  balance: "",
  storage: "",
  blocknumber: "",
};

export const EthState = {
  encode(message: EthState, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.slot !== "") {
      writer.uint32(26).string(message.slot);
    }
    if (message.balance !== "") {
      writer.uint32(34).string(message.balance);
    }
    if (message.storage !== "") {
      writer.uint32(42).string(message.storage);
    }
    if (message.blocknumber !== "") {
      writer.uint32(50).string(message.blocknumber);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EthState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEthState } as EthState;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.slot = reader.string();
          break;
        case 4:
          message.balance = reader.string();
          break;
        case 5:
          message.storage = reader.string();
          break;
        case 6:
          message.blocknumber = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EthState {
    const message = { ...baseEthState } as EthState;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.slot !== undefined && object.slot !== null) {
      message.slot = String(object.slot);
    } else {
      message.slot = "";
    }
    if (object.balance !== undefined && object.balance !== null) {
      message.balance = String(object.balance);
    } else {
      message.balance = "";
    }
    if (object.storage !== undefined && object.storage !== null) {
      message.storage = String(object.storage);
    } else {
      message.storage = "";
    }
    if (object.blocknumber !== undefined && object.blocknumber !== null) {
      message.blocknumber = String(object.blocknumber);
    } else {
      message.blocknumber = "";
    }
    return message;
  },

  toJSON(message: EthState): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.address !== undefined && (obj.address = message.address);
    message.slot !== undefined && (obj.slot = message.slot);
    message.balance !== undefined && (obj.balance = message.balance);
    message.storage !== undefined && (obj.storage = message.storage);
    message.blocknumber !== undefined &&
      (obj.blocknumber = message.blocknumber);
    return obj;
  },

  fromPartial(object: DeepPartial<EthState>): EthState {
    const message = { ...baseEthState } as EthState;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.slot !== undefined && object.slot !== null) {
      message.slot = object.slot;
    } else {
      message.slot = "";
    }
    if (object.balance !== undefined && object.balance !== null) {
      message.balance = object.balance;
    } else {
      message.balance = "";
    }
    if (object.storage !== undefined && object.storage !== null) {
      message.storage = object.storage;
    } else {
      message.storage = "";
    }
    if (object.blocknumber !== undefined && object.blocknumber !== null) {
      message.blocknumber = object.blocknumber;
    } else {
      message.blocknumber = "";
    }
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
