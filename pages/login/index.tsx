import Router from 'next/router'

import { toast } from 'react-hot-toast'
import { GoogleLoginButton } from 'react-social-login-buttons';

import { firebaseAuth, db, googleProvider } from '../../lib/firebase'
import { signInWithPopup } from 'firebase/auth';
import { doc, getDoc, setDoc } from "firebase/firestore";

import styles from "../../styles/Login.module.css"

export default function LoginPage() {
    const signInWithGoogle = async () => {
        signInWithPopup(firebaseAuth, googleProvider)
            .then( async (res) => {
                const user = res.user;

                const docSnap = await getDoc(doc(db, "Users", user.uid));
                if (!docSnap.exists()) {
                    await setDoc(doc(db, "Users", user.uid), {
                        uid: user.uid,
                        photoURL: user.photoURL,
                        displayName: user.displayName,
                    });
                }

                Router.push("/profile");
                toast.success(`Logged in as ${user.displayName}`)
            }).catch( (err) => {
                toast.error(err.message);
            });
    }

    return (
        <div className={styles.loginContainer}>
            <div className={styles.buttonContainer}>
                <GoogleLoginButton onClick={signInWithGoogle}/>
            </div>
        </div>
    );
}