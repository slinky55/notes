
import { useForm } from 'react-hook-form'
import Router from 'next/router'

import { toast } from 'react-hot-toast'

import { pocketbase } from '../../lib/backend'

import styles from "../../styles/Login.module.css"

export default function LoginPage() {
    const { register, handleSubmit, watch, formState: {errors}} = useForm();

    const onSubmit = async (data: any) => {
        try {
            const userData = await pocketbase.users.authViaEmail(data.email, data.password);

            console.log(userData);

            toast.success('Logged in!');
            Router.push("/profile");
        } catch (err) {
            toast.error("Failed to log in");
        }
    }

    return (
        <div className={styles.loginContainer}>
            <form onSubmit={handleSubmit(onSubmit)}>
                <div className={styles.inputContainer}>
                    <label>Email</label>
                    <input type="text" placeholder="email@example.com" {...register("email", { required: true })} />
                    {errors.email ? <span>This field is required</span> : (<span></span>) }
                </div>
                <div className={styles.inputContainer}>
                    <label>Password</label>
                    <input type="password" {...register("password", { required: true })} />
                    {errors.password ? <span>This field is required</span> : (<span></span>) }
                </div>
                <input type="submit" value="Login"/>
            </form> 
        </div>
    );
}