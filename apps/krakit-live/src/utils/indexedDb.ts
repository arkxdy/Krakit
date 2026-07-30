import { openDB } from "idb";

const DB_NAME = "my-app-db";
const STORE_NAME = "my-store";

export async function getDb() {
  return openDB(DB_NAME, 1, {
    upgrade(db) {
      if (!db.objectStoreNames.contains(STORE_NAME)) {
        db.createObjectStore(STORE_NAME);
      }
    },
  });
}

export async function setItem(key: string, value: any) {
  const db = await getDb();
  return db.put(STORE_NAME, value, key);
}

export async function getItem(key: string) {
  const db = await getDb();
  return db.get(STORE_NAME, key);
}

export async function deleteItem(key: string) {
  const db = await getDb();
  return db.delete(STORE_NAME, key);
}
