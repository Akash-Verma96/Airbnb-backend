import { AsyncLocalStorage } from "node:async_hooks";

type asyncloaclstorageType = {
    correlationid : string;
}


export const asyncLocalStorage = new AsyncLocalStorage<asyncloaclstorageType>();


export const getCorrelationId = () => {
    const asyncStore = asyncLocalStorage.getStore();
    return asyncStore?.correlationid || 'unknown-error-while-creating-correlationid';
}