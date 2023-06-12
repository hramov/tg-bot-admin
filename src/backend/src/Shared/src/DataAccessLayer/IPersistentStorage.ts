export interface PersistentStorage {
	getStorage: <T>(type: new () => T) => T;
}
