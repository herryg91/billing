import Head from 'next/head'
import { useAuth } from '@/pkg/hook/useAuth';
import { useState } from 'react';
import { AuthApi } from '@/repositories/auth-api';
import { ApiErrorResponse } from '@/pkg/api/response';
import { toast } from 'react-toastify';
import * as yup from 'yup'
import { useForm } from 'react-hook-form'
import { yupResolver } from '@hookform/resolvers/yup'
import { Button, Loading } from 'react-daisyui';

export default function LoginPage() {
  const { login } = useAuth()
  const [isLoading, setIsLoading] =  useState(false)

  const ValidationSchema: yup.ObjectSchema<{email: string, password: string}> = yup.object().shape({
    email: yup.string().email().required(),
    password: yup.string().required(),
  })
  
  const { register, handleSubmit, formState: { errors }, setError } = useForm<{email: string, password: string}>({ mode: 'all', resolver: yupResolver(ValidationSchema) });

  const onSubmitHandler = async (data: {email: string, password: string}) => {
    try {
      setIsLoading(true)
      const resp = await AuthApi.Login(data.email, data.password)
      toast.success("Login success")
      login(resp.data)
    } catch (error) {
      const api_err = error as ApiErrorResponse
      api_err.other_errors.forEach( (e) => {
        switch(e.field){
          case "email" : setError("email", { type: "focus", message: e.message }, { shouldFocus: true }); break;
          case "password": setError("password", { type: "focus", message: e.message }, { shouldFocus: true }); break;
      }
      })
      toast.error((error as ApiErrorResponse).message)
    } finally {
      setIsLoading(false)
    }
  };


  return <>
    <Head>
      <title>Sign in to dashboard - Blogstraps</title>
      <meta name="robots" content="noindex,nofollow" />
      <meta name='description' content='Create and embed blogs in minutes through our blog management system' />
      <meta property="og:title" content={"Sign in to dashboard - Blogstraps"} />
      <meta property="og:description" content={"Create and embed blogs in minutes through our blog management system"} />
      <meta property="og:url" content={"https://blogstraps.com"} />
      <meta property="og:type" content="website" />
      <meta name="twitter:card" content="summary_large_image" />
      <meta name="twitter:domain" content="blogstraps.com" />
      <meta name="twitter:title" content={"Sign in to dashboard - Blogstraps"} />
      <meta name="twitter:description" content={"Create and embed blogs in minutes through our blog management system"} />
    </Head>
    
    <div className="relative flex flex-col justify-center h-screen overflow-hidden">
      <div className="w-full p-6 m-auto bg-white rounded-3xl shadow-md lg:max-w-lg">
        <h1 className="text-2xl mt-4 font-semibold text-center mb-2">Sign In to <b>billing.com</b></h1>
        <form className="space-y-4 mt-4" onSubmit={ handleSubmit(onSubmitHandler) }>
          <div>
              <label className="label">
                  <span className="text-base label-text">Email</span>
              </label>
              <input {...register("email")} type="text" placeholder="Email Address" className="w-full input input-bordered " />
              <p className="mt-2 text-sm text-red-600 dark:text-red-500">{(errors?.email && <>{errors.email.message}</>)}</p>
          </div>
          <div>
              <label className="label">
                  <span className="text-base label-text">Password</span>
              </label>
              <input {...register("password")} type="password" placeholder="Enter Password"
                  className="w-full input input-bordered " />
              <p className="mt-2 text-sm text-red-600 dark:text-red-500">{(errors?.password && <>{errors.password.message}</>)}</p>
          </div>
          <div>
              <Button type="submit" disabled={isLoading} color="success" className='w-full text-white'>LOGIN {isLoading && <Loading />}</Button>
          </div>  
        </form>
      </div>
    </div>
  </>
}