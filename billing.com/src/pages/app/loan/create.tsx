import DashboardLayout from '@/components/layouts/dashboard-layout';
import Head from 'next/head';
import { NextPage } from 'next';
import { Button, Input, Loading, Table } from 'react-daisyui';
import * as yup from 'yup'
import { useForm } from "react-hook-form";
import { yupResolver } from '@hookform/resolvers/yup'
import { useState } from 'react';
import { useFetch } from '@/pkg/hook/useFetch';
import { LoanApi } from '@/repositories/loan-api';
import { LoanSimulation } from '@/repositories/loan-api/entity';
import { ApiErrorResponse } from '@/pkg/api/response';
import { toast } from 'react-toastify';
import { useRouter } from 'next/router';

const LoanCreatePage: NextPage<void> = () => {
  const router = useRouter()
  const [submitting, setSubmitting] = useState(false)
  const simulateLoan = useFetch<LoanSimulation>(LoanApi.SimulateLoan)

  const ValidationSchema = yup.object().shape({
    description: yup.string().required(),
    installment_length: yup.number().required(),
    principal: yup.number().required(),
  })
  const { register, handleSubmit, formState: { errors }, setError, trigger, getValues } = useForm<{description: string, installment_length: number,  principal: number}>({ mode: 'onChange', resolver: yupResolver(ValidationSchema), defaultValues: {
    description: "",
    installment_length: 50,
    principal: 5000000,
  }});

  const onSubmitHandler = async (data: {description: string, principal: number, installment_length: number, }) => {
      try {
        setSubmitting(true);
        await LoanApi.CreateLoan(data.description, data.principal, data.installment_length)
        toast.success("Pinjaman berhasil dibuat");
        router.push("/app/loan")
      } catch (error) {
        if (error as ApiErrorResponse) {
          const err_api=(error as ApiErrorResponse)
          if(err_api.other_errors.length > 0){
            err_api.other_errors.forEach((e) => {
              switch(e.field){
                case "principal" : setError("principal", { type: "focus", message: e.message }, { shouldFocus: true }); break;
                case "description" : setError("description", { type: "focus", message: e.message }, { shouldFocus: true }); break;
                case "installment_length" : setError("installment_length", { type: "focus", message: e.message }, { shouldFocus: true }); break;
              }
            })
          }
          toast.error(err_api.message);
        } else {
            console.log("Unknown error:", error);
            toast.error("Internal Error");
        }
      } finally {
        setSubmitting(false);
      }
  }

  const doSimulateLoan = async () => {
    setSubmitting(true);
    const ok = await trigger()
    if (ok){
      simulateLoan.request(getValues('principal'), getValues('installment_length'))
    }
    setSubmitting(false);
  }

  return <DashboardLayout title='Buat Pengajuan Pinjaman' className='p-8'>
    <Head>
      <title>Buat Pengajuan Pinjaman</title>
      <meta name="robots" content="noindex,nofollow" />
    </Head>

    <div className='grid md:grid-cols-2  gap-6'>
      <div className='w-full bg-white border rounded-xl p-6 max-h-fit'>
        <h1 className='text-center font-bold mb-8'>Pengajuan Pinjaman</h1>
        <form className="flex flex-col gap-y-4 mx-auto" onSubmit={handleSubmit(onSubmitHandler)} >
          <div className='relative'>
            <div className='absolute ml-3 pt-1.5 text-gray-500 text-xs font-light'>Deskripsi Pinjaman</div>
            <Input {...register("description")} className='h-12 pt-4 w-full' size='sm' />
            {(errors?.description && <p className="mt-2 text-sm text-red-600 dark:text-red-500">{errors.description.message}</p>) }
          </div>
          <div className='grid grid-cols-2 gap-x-4'>
            <div className='relative'>
              <div className='absolute ml-3 pt-1.5 text-gray-500 text-xs font-light'>Jumlah Pinjaman</div>
              <Input {...register("principal")} type="number" min={1} className='h-12 pt-4 w-full' size='sm' />
              {(errors?.principal && <p className="mt-2 text-sm text-red-600 dark:text-red-500">{errors.principal.message}</p>) }
            </div>
            <div className='relative'>
              <div className='absolute ml-3 pt-1.5 text-gray-500 text-xs font-light'>Tenor (1 Tahun = 50 Minggu)</div>
              <Input {...register("installment_length")} type="number" min={1} className='h-12 pt-4 w-full' size='sm' />
              {(errors?.installment_length && <p className="mt-2 text-sm text-red-600 dark:text-red-500">{errors.installment_length.message}</p>) }
            </div>
            
          </div>
          
          <div className='relative'>
            <div className='absolute ml-3 pt-1.5 text-gray-500 text-xs font-light'>Bunga</div>
            <Input className='h-12 pt-4 w-full' size='sm' value="10% Flat per Annum" disabled />
          </div>

            <Button type='button' color='ghost' size='md' disabled={submitting} onClick={doSimulateLoan}>Hitung Simulasi</Button>
            <Button type='submit' color='success' className="text-white" size='md' disabled={submitting}>Buat Pengajuan {submitting && <Loading size="sm" />}</Button>
        </form>
      </div>
      {
      simulateLoan.data &&
      <div>
        <h3 className='font-bold text-center mb-4'> Simulasi Pinjaman</h3>
        <div className='w-full bg-white border rounded-xl p-6 max-h-fit'>
          <Table>
            <Table.Body>
              <Table.Row>
                <span>Tenor</span>
                <span>{ simulateLoan.data.installment_length } Minggu</span>
              </Table.Row>
              <Table.Row>
                <span>Pokok</span>
                <span>{ simulateLoan.data.principal.toLocaleString("id") }</span>
              </Table.Row>
              <Table.Row>
                <span>Bunga</span>
                <span>{ simulateLoan.data.interest_amount.toLocaleString("id") }</span>
              </Table.Row>
              <Table.Row>
                <span>Total Pinjaman</span>
                <span>{ simulateLoan.data.total_amount.toLocaleString("id") }</span>
              </Table.Row>
            </Table.Body>
          </Table>
        </div>
        <div className='mt-4'>
          <Table>
            <Table.Head>
              <span>Cicilan Ke-</span>
              <span>Pokok</span>
              <span>Bunga</span>
              <span>Total</span>
            </Table.Head>
            <Table.Body>
              {
              simulateLoan.data.billings.map((l,index) => {
              return <Table.Row key={`billing-${index}`}>
                <span>{l.installment_number}</span>
                <span>{l.principal.toLocaleString("id")}</span>
                <span>{l.interest_amount.toLocaleString("id")}</span>
                <span>{l.total_amount.toLocaleString("id")}</span>
              </Table.Row>
              })
            }
            </Table.Body>
          </Table>
          </div>
      </div>
      }
    </div>

    
  </DashboardLayout>
}

export default LoanCreatePage

