import DashboardLayout from '@/components/layouts/dashboard-layout';
import Head from 'next/head';
import { GetServerSideProps, NextPage } from 'next';
import LogoBCA from '@/assets/images/BCA.svg';
import Image from 'next/image';
import { Button, Link, Skeleton } from 'react-daisyui';
import { useEffect, useState } from 'react';
import { toast } from 'react-toastify';
import { useRouter } from 'next/router';
import { ApiErrorResponse } from '@/pkg/api/response';
import { BillingApi } from '@/repositories/billing-api';
import { Billing } from '@/repositories/billing-api/entity';
import { useFetch } from '@/pkg/hook/useFetch';
import moment from 'moment';
import { HiCheck } from 'react-icons/hi';

export const getServerSideProps: GetServerSideProps<{code: string, installment: number, refUrl: string}> = async (context) => {
  const code = context.params?.code?.toString() ?? ""
  const installment = Number(context.params?.installment)
  const ref = context.query.ref?.toString()
  return {
    props: { 
      code: code,
      installment: installment,
      refUrl:ref ?? ""
     }, 
  }
}

const LoanBillingPaymentPage: NextPage<{code: string, installment: number, refUrl: string}> = (param) => {
  const router = useRouter()
  const [submitting, setSubmitting] = useState(false)

  const generatePaymentInfo = useFetch<Billing>(BillingApi.GenerateBillingPayment)
  useEffect(() => {
    generatePaymentInfo.request(param.code, param.installment)
  },[])

  const settlePayment = async () => {
    try {
      setSubmitting(true)
      await BillingApi.SettleBillingPayment(param.code, param.installment)
      toast.success("Terima kasih telah mengkonfirmasi pembayaran")
      router.push(`/app/loan/${param.code}`)
    } catch (error) {
      toast.error((error as ApiErrorResponse).message)
    } finally {
      setSubmitting(false)
    }
  }

  return <DashboardLayout title='Loan Payment' className='p-8'>
    <Head>
      <title>Loan Payment</title>
      <meta name="robots" content="noindex,nofollow" />
    </Head>
   
    <div className='mx-auto p-6 w-full max-w-sm border rounded-xl mt-8'>
      <div className='flex flex-col gap-y-4'>
        <div className='flex justify-between gap-4'> 
          <div>No Pinjaman </div>
          <div>{generatePaymentInfo.data?.loan_code || <Skeleton />}</div>
        </div>
        <div className='flex justify-between gap-4'> 
          <div>Cicilan Ke</div>
          <div>{generatePaymentInfo.data?.installment_number || <Skeleton />}</div>
        </div>
        <div className='text-center'>
          <div className='text-sm'>Total Pembayaran</div>
          <div className='font-bold text-2xl'>{(generatePaymentInfo.data?.total_amount ?? 0).toLocaleString("id")}</div>
        </div>
        <div className='border rounded-xl p-4 flex gap-x-4'>
          <Image src={LogoBCA} alt='bca' height={18}/>
          <div className='text-sm'>
            <div className='font-bold'>{generatePaymentInfo.data?.payment_va}</div>
            <div>A/N billing.com</div>
          </div>
        </div>
        {generatePaymentInfo.data &&
        <div className='text-xs'>Lakukan Pembayaran Sebelum: { moment(generatePaymentInfo.data.payment_expired_at).format("DD-MM-YYYY, HH:mm:ss") }</div>
        }
        {
          generatePaymentInfo.data?.payment_status !== 'PAID' ?
          <Button type="button" disabled={submitting} color="success" className='text-white' onClick={settlePayment}>Konfirmasi Pembayaran</Button>:
          <div className='text-xs flex gap-x-1 items-center text-success font-bold'><HiCheck/> Sudah Dibayar</div>
        }
        <Link href={param.refUrl ==="/app" ? "/app" :  `/app/loan/${param.code}`}><Button type="button" color="ghost" className='w-full'>Kembali</Button></Link>
      </div>
      
    </div>

  </DashboardLayout>
}

export default LoanBillingPaymentPage

