import { initializeApp } from "firebase/app";
import { getAuth, GoogleAuthProvider } from "firebase/auth";
import { getFirestore } from "firebase/firestore";

const firebaseConfig = {
    apiKey: "AIzaSyCB1Shvh--mOUvvT8PQeIXgcnY6wUm2E7Y",
    authDomain: "simplnotes.firebaseapp.com",
    projectId: "simplnotes",
    storageBucket: "simplnotes.appspot.com",
    messagingSenderId: "1021226269602",
    appId: "1:1021226269602:web:868c74ee746422cba450fb",
    measurementId: "G-C8D7103RWE"
};

export const firebase = initializeApp(firebaseConfig);
export const firebaseAuth = getAuth(firebase);
export const db = getFirestore(firebase);

export const googleProvider = new GoogleAuthProvider();
  