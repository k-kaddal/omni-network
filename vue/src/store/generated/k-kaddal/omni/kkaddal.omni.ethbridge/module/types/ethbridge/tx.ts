/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { State } from "../ethbridge/state";

export const protobufPackage = "kkaddal.omni.ethbridge";

export interface MsgStoreState {
  creator: string;
  address: string;
  slot: string;
}

export interface MsgStoreStateResponse {
  state: State | undefined;
}

const baseMsgStoreState: object = { creator: "", address: "", slot: "" };

export const MsgStoreState = {
  encode(message: MsgStoreState, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.slot !== "") {
      writer.uint32(26).string(message.slot);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgStoreState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgStoreState } as MsgStoreState;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.slot = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgStoreState {
    const message = { ...baseMsgStoreState } as MsgStoreState;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
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
    return message;
  },

  toJSON(message: MsgStoreState): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.address !== undefined && (obj.address = message.address);
    message.slot !== undefined && (obj.slot = message.slot);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgStoreState>): MsgStoreState {
    const message = { ...baseMsgStoreState } as MsgStoreState;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
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
    return message;
  },
};

const baseMsgStoreStateResponse: object = {};

export const MsgStoreStateResponse = {
  encode(
    message: MsgStoreStateResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.state !== undefined) {
      State.encode(message.state, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgStoreStateResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgStoreStateResponse } as MsgStoreStateResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.state = State.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgStoreStateResponse {
    const message = { ...baseMsgStoreStateResponse } as MsgStoreStateResponse;
    if (object.state !== undefined && object.state !== null) {
      message.state = State.fromJSON(object.state);
    } else {
      message.state = undefined;
    }
    return message;
  },

  toJSON(message: MsgStoreStateResponse): unknown {
    const obj: any = {};
    message.state !== undefined &&
      (obj.state = message.state ? State.toJSON(message.state) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgStoreStateResponse>
  ): MsgStoreStateResponse {
    const message = { ...baseMsgStoreStateResponse } as MsgStoreStateResponse;
    if (object.state !== undefined && object.state !== null) {
      message.state = State.fromPartial(object.state);
    } else {
      message.state = undefined;
    }
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  StoreState(request: MsgStoreState): Promise<MsgStoreStateResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  StoreState(request: MsgStoreState): Promise<MsgStoreStateResponse> {
    const data = MsgStoreState.encode(request).finish();
    const promise = this.rpc.request(
      "kkaddal.omni.ethbridge.Msg",
      "StoreState",
      data
    );
    return promise.then((data) =>
      MsgStoreStateResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
