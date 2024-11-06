import DashboardLayout from '@/components/layouts/dashboard-layout';
import Head from 'next/head';
import { NextPage } from 'next';
import { Badge, Button, Link, Skeleton, Stats, Table } from 'react-daisyui';
import { HiOutlineBadgeCheck, HiOutlineCreditCard, HiOutlinePresentationChartLine, HiOutlineXCircle } from 'react-icons/hi';
import { useFetch } from '@/pkg/hook/useFetch';
import { LoanApi } from '@/repositories/loan-api';
import { useEffect } from 'react';
import { useUser } from '@/pkg/hook/useUser';
import moment from 'moment';
import { BillingApi } from '@/repositories/billing-api';
import { Billings } from '@/repositories/billing-api/entity';

const DashboardIndexPage: NextPage<void> = () => {
  const userCtx = useUser()
  const getMySummary = useFetch<{loan_count: number;total_outstanding: number;is_delinquent: boolean;}>(LoanApi.GetMySummary)
  const getBillingOverDue = useFetch<Billings>(BillingApi.GetBillingOverDue)

  useEffect(() => {
    getMySummary.request()
    getBillingOverDue.request()
  }, [])

  return <DashboardLayout title='Dashboard' className='p-8'>
    <Head>
      <title>Dashboard</title>
      <meta name="robots" content="noindex,nofollow" />
    </Head>
    <h1 className='font-bold text-lg mb-8'>Selamat Datang, {userCtx.user?.name}</h1>
    <div className='grid grid-cols-3 gap-8'>
      <Stats className="shadow">
        <Stats.Stat>
          <Stats.Stat.Figure className="text-neutral">
            <HiOutlineCreditCard size={32}/>
          </Stats.Stat.Figure>
          <Stats.Stat.Title className='font-light'>Jumlah Pinjaman Aktif</Stats.Stat.Title>
          <Stats.Stat.Value>{ getMySummary.data?.loan_count }</Stats.Stat.Value>
        </Stats.Stat>
      </Stats>
      <Stats className="shadow">
        <Stats.Stat>
          <Stats.Stat.Figure className="text-neutral">
            <HiOutlinePresentationChartLine size={32}/>
          </Stats.Stat.Figure>
          <Stats.Stat.Title className='font-light'>Total Outstanding</Stats.Stat.Title>
          <Stats.Stat.Value>{(getMySummary.data?.total_outstanding ?? 0).toLocaleString("id")}</Stats.Stat.Value>
        </Stats.Stat>
      </Stats>
      <Stats className="shadow">
        <Stats.Stat>
          <Stats.Stat.Figure className={`${!getMySummary.data ? "text-success" :  getMySummary.data.is_delinquent ? "text-error" : "text-success" }`}>
            {
              (getMySummary.data && getMySummary.data.is_delinquent )?
              <HiOutlineXCircle  size={32}/>:
              <HiOutlineBadgeCheck size={32}/>
            }
          </Stats.Stat.Figure>
          <Stats.Stat.Title className='font-light'>Credit Score</Stats.Stat.Title>
          <Stats.Stat.Value className={`${!getMySummary.data ? "text-success" :  getMySummary.data.is_delinquent ? "text-error" : "text-success" }`}>
            { !getMySummary.data ? <Skeleton/> :
              getMySummary.data.is_delinquent ? "Buruk" :
              "Baik" 
            }</Stats.Stat.Value>
            {getMySummary.data && getMySummary.data.is_delinquent &&
              <Stats.Stat.Desc>Mohon lunasi tunggakan anda</Stats.Stat.Desc> 
            }
        </Stats.Stat>
      </Stats>
    </div>
    
    <h2 className='font-bold text-lg my-8'>Daftar Jatuh Tempo</h2>
    {
      getBillingOverDue.data && getBillingOverDue.data.billings.length === 0 &&
      <div>Tidak ada pinjaman mendekati atau melewati jatuh tempo</div>
    }
    {
      getBillingOverDue.data && getBillingOverDue.data.billings.length > 0 &&
      <div className="overflow-x-auto">
      <Table className="rounded-box">
        <Table.Head>
          <span>No Pinjaman</span>
          <span>Cicilan Ke-</span>
          <span>Tanggal Jatuh Tempo</span>
          <span>Jumlah</span>
          <span>Status</span>
          <span></span>
        </Table.Head>
        <Table.Body>
          {
            getBillingOverDue.data.billings.map((b, index) => {
              return <Table.Row key={`due-billing-${index}`}>
              <div>{b.loan_code}</div>
              <div>{b.installment_number}</div>
              <div>{moment(b.due_date).format("DD-MM-YYYY")}</div>
              <div>{b.total_amount.toLocaleString('id')}</div>
              <div>{ 
              b.payment_status === "PAID" ? 
              <Badge color="success" className='text-white' size="sm">Sudah Dibayar</Badge> : 
              b.payment_status === "WAIT_FOR_PAYMENT" ? 
              <Badge color="warning" className='text-white' size="sm">Menunggu Pembayaran/Verifikasi</Badge> : 
              moment(b.due_date).format("YYYY-MM-DD") < moment(new Date()).format("YYYY-MM-DD")?
              <Badge color="error" className='text-white' size="sm">Lewat Jatuh Tempo</Badge> :
              <Badge color="neutral" className='text-white' size="sm">Akan Jatuh Tempo</Badge> 
              }
              </div>
              <div>
                <Link href={`/app/loan/${b.loan_code}/${b.installment_number}?ref=/app`}>
                  <Button type="button" size="sm" color="success" className='text-white'>Proses Pembayaran</Button>
                </Link>
              </div>
            </Table.Row>
            })
          }
        </Table.Body>
      </Table>
      </div>
    }
  </DashboardLayout>
}

export default DashboardIndexPage

