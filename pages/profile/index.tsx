import { collection, getDocs, query } from "firebase/firestore";
import { firebaseAuth, db } from "../../lib/firebase";

export default function Profile() {
    
    const getNotes = async () => {
        const user = firebaseAuth.currentUser;

        console.log(user);
    }

    return (
        <>
            <h1>Profile</h1>
            {getNotes}
        </>
    )
}