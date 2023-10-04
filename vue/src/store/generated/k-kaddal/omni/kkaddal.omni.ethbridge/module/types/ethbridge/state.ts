/* eslint-disable */
import { Metadata } from "../ethbridge/metadata";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "kkaddal.omni.ethbridge";

export interface State {
  address: string;
  slot: string;
  balance: string;
  data: string;
  metadata: Metadata | undefined;
}

const baseState: object = { address: "", slot: "", balance: "", data: "" };

export const State = {
  encode(message: State, writer: Writer = Writer.create()): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.slot !== "") {
      writer.uint32(18).string(message.slot);
    }
    if (message.balance !== "") {
      writer.uint32(26).string(message.balance);
    }
    if (message.data !== "") {
      writer.uint32(34).string(message.data);
    }
    if (message.metadata !== undefined) {
      Metadata.encode(message.metadata, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): State {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseState } as State;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 2:
          message.slot = reader.string();
          break;
        case 3:
          message.balance = reader.string();
          break;
        case 4:
          message.data = reader.string();
          break;
        case 5:
          message.metadata = Metadata.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): State {
    const message = { ...baseState } as State;
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
    if (object.data !== undefined && object.data !== null) {
      message.data = String(object.data);
    } else {
      message.data = "";
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      message.metadata = Metadata.fromJSON(object.metadata);
    } else {
      message.metadata = undefined;
    }
    return message;
  },

  toJSON(message: State): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.slot !== undefined && (obj.slot = message.slot);
    message.balance !== undefined && (obj.balance = message.balance);
    message.data !== undefined && (obj.data = message.data);
    message.metadata !== undefined &&
      (obj.metadata = message.metadata
        ? Metadata.toJSON(message.metadata)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<State>): State {
    const message = { ...baseState } as State;
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
    if (object.data !== undefined && object.data !== null) {
      message.data = object.data;
    } else {
      message.data = "";
    }
    if (object.metadata !== undefined && object.metadata !== null) {
      message.metadata = Metadata.fromPartial(object.metadata);
    } else {
      message.metadata = undefined;
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
