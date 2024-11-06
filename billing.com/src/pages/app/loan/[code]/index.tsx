import DashboardLayout from '@/components/layouts/dashboard-layout';
import Head from 'next/head';
import { GetServerSideProps, NextPage } from 'next';
import { Badge, Button, Link, Table } from 'react-daisyui';
import { useFetch } from '@/pkg/hook/useFetch';
import { LoanApi } from '@/repositories/loan-api';
import { Loan } from '@/repositories/loan-api/entity';
import { useEffect } from 'react';
import { BillingApi } from '@/repositories/billing-api';
import { Billings } from '@/repositories/billing-api/entity';
import moment from 'moment';

export const getServerSideProps: GetServerSideProps<{code: string}> = async (context) => {
  const code = context.params?.code?.toString() ?? ""
  return {
    props: { 
      code: code
     }, 
  }
}

const LoanDetailPage: NextPage<{code: string}> = (param) => {
  const getLoanDetail = useFetch<Loan>(LoanApi.GetLoanByCode)
  const getBillings = useFetch<Billings>(BillingApi.GetBillingByLoanCode)

  useEffect(() => {
    getLoanDetail.request(param.code)
    getBillings.request(param.code)
  }, [])

  return <DashboardLayout title='Loan Detail' className='p-8'>
    <Head>
      <title>Loan Detail</title>
      <meta name="robots" content="noindex,nofollow" />
    </Head>

    {
    getLoanDetail.data &&
    <div className='w-full bg-white border rounded-xl p-6 max-h-fit'>
      <Table>
        <Table.Body>
          <Table.Row>
            <span>Tenor</span>
            <span>{ getLoanDetail.data.installment_length } Minggu </span>
          </Table.Row>
          <Table.Row>
            <span>Pokok</span>
            <span>{ getLoanDetail.data.principal.toLocaleString("id") }</span>
          </Table.Row>
          <Table.Row>
            <span>Bunga</span>
            <span>{ getLoanDetail.data.interest_amount.toLocaleString("id") }</span>
          </Table.Row>
          <Table.Row>
            <span>Total Pinjaman</span>
            <span>{ getLoanDetail.data.total_amount.toLocaleString("id") }</span>
          </Table.Row>
        </Table.Body>
      </Table>
    </div>
    }
    {
    getBillings.data &&
    <div className='mt-4'>
      <Table>
        <Table.Head>
          <span>Cicilan Ke-</span>
          <span>Jatuh Tempo</span>
          <span>Pokok</span>
          <span>Bunga</span>
          <span>Total</span>
          <span>Status</span>
          <span></span>
        </Table.Head>
        <Table.Body>
          {
          getBillings.data.billings.map((l,index) => {
          return <Table.Row key={`billing-${index}`}>
            <span>{l.installment_number}</span>
            <span>{moment(l.due_date).format("DD-MM-YYYY")}</span>
            <span>{l.principal.toLocaleString("id")}</span>
            <span>{l.interest_amount.toLocaleString("id")}</span>
            <span>{l.total_amount.toLocaleString("id")}</span>
            <div className=''>{ 
              l.payment_status === "PAID" ? 
              <Badge color="success" className='text-white' size="sm">Sudah Dibayar</Badge> : 
              l.payment_status === "WAIT_FOR_PAYMENT" ? 
              <Badge color="warning" className='text-white' size="sm">Menunggu Pembayaran/Verifikasi</Badge> : 
              moment(l.due_date).format("YYYY-MM-DD") < moment(new Date()).format("YYYY-MM-DD")?
              <Badge color="error" className='text-white' size="sm">Lewat Jatuh Tempo</Badge> :
              <Badge color="neutral" className='text-white' size="sm">Belum Dibayar</Badge> 
              }
            </div>
            <div>
              { l.payment_status !== "PAID" &&
              <Link href={`/app/loan/${param.code}/${l.installment_number}`}><Button type="button" color="success" className='text-white' size="sm">Proses Pembayaran</Button></Link>
              }
            </div>
          </Table.Row>
          })
        }
        </Table.Body>
      </Table>
    </div>
    }
   
    {
      getLoanDetail.data &&
      <div>
        
        
      </div>
      }

  </DashboardLayout>
}

export default LoanDetailPage

