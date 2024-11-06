import DashboardLayout from '@/components/layouts/dashboard-layout';
import Head from 'next/head';
import { NextPage } from 'next';
import { Badge, Button, Table } from 'react-daisyui';
import { HiEye, HiPlus } from 'react-icons/hi';
import { useRouter } from 'next/router';
import { useFetch } from '@/pkg/hook/useFetch';
import { LoanApi } from '@/repositories/loan-api';
import { Loans } from '@/repositories/loan-api/entity';
import { useEffect } from 'react';

const LoanIndexPage: NextPage<void> = () => {
  const router = useRouter()
  const getLoans = useFetch<Loans>(LoanApi.GetLoans)

  useEffect(() => {
    getLoans.request()
  }, [])

  return <DashboardLayout title='My Loan' className='p-8'>
    <Head>
      <title>My Loan</title>
      <meta name="robots" content="noindex,nofollow" />
    </Head>
   
    <div className='flex justify-between items-center my-4'>
      <div className='mb-4 flex justify-end'>
          <Button type="button" color="success" className={`text-white`} onClick={() => router.push("/app/loan/create")}><HiPlus /> Pengajuan Pinjaman</Button>
      </div>
    </div>

    <div className="overflow-x-auto">
      <Table className="rounded-box">
        <Table.Head>
          <span>No Pinjaman</span>
          <span>Deskripsi</span>
          <span>Tenor</span>
          <span>Total Pinjaman</span>
          <span>Outstanding</span>
          <span>Status</span>
          <span></span>
        </Table.Head>
        <Table.Body>
        {getLoans.data && getLoans.data.loans.length === 0 &&
          <tr>
            <td colSpan={6} className='text-center'>Belum ada pinjaman</td>
          </tr>
        }
        {
          getLoans.data?.loans.map((l, index) => {
          return <Table.Row key={`loan-${index}`} onClick={() => router.push(`/app/loan/${l.code}`)} hover>
            <span>{l.code}</span>
            <span>{l.description}</span>
            <span>{l.installment_length} Minggu</span>
            <span>{(l.total_amount).toLocaleString("id")}</span>
            <span>{(l.outstanding).toLocaleString("id")}</span>
            <span>{
              l.status === "PENDING" ? <Badge color="neutral" className='text-white'>Menunggu Approval</Badge> :
              l.status === "APPROVED" ? <Badge color="warning" className='text-white'>Pencairan</Badge> :
              l.status === "ACTIVE" ? <Badge color="warning" className='text-white'>Aktif</Badge> :
              l.status === "DONE" ? <Badge color="success" className='text-white'>Selesai</Badge> :
              <></>}
            </span>
            <div>
              <Button type="button" shape='circle' color="success" className='text-white' size='sm' onClick={() => router.push(`/app/loan/${l.code}`)}><HiEye /></Button>
            </div>
          </Table.Row>
          })
        }
        </Table.Body>
      </Table>
    </div>

  </DashboardLayout>
}

export default LoanIndexPage

