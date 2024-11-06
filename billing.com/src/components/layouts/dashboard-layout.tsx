import React, { ReactNode, useState } from "react";
import { FiMenu } from "react-icons/fi";
import Link from "next/link";
import { Button, Drawer, Dropdown, Menu, Navbar } from "react-daisyui";
import { HiOutlineChartBar, HiOutlineCreditCard } from "react-icons/hi";
import Avatar from "react-avatar";
import LogoRect from "@/assets/images/logo-rect.svg";
import Image from "next/image";
import { useAuth } from "@/pkg/hook/useAuth";
import { useUser } from "@/pkg/hook/useUser";


const DashboardLayout = (props :{children? :ReactNode, title?: string, className?: string }) => {
  const { logout } = useAuth()
  const currentUser = useUser()

  const [drawer_open, set_drawer_open] = useState(false)

  return (
    <div className="w-full lg:pl-72">
      <NavigationSidebar className="fixed start-0 z-30 hidden lg:block" logout={logout} />
      <Drawer
        open={drawer_open}
        onClickOverlay={() => set_drawer_open(!drawer_open)}
        side={<NavigationSidebar className="block" logout={() => {}} />}
        className="min-h-screen"
      >
          <Navbar className="px-2 font-sans border-b border-gray-300" >
          <Navbar.Start>
            <div className="flex-none lg:hidden">
              <Button shape="square" color="ghost" onClick={() => set_drawer_open(!drawer_open)}>
                <FiMenu className="inline-block w-6 h-6 stroke-current" />
              </Button>
            </div>
            <div className="font-bold ml-2">{props.title}</div>
          </Navbar.Start>
          <Navbar.End>
            <Dropdown end color="accent">
              <Button tag="label" tabIndex={0} color="ghost" className="avatar" shape="circle">
                <Avatar name={currentUser.user?.name ?? ""} size="40" round={true} />
              </Button>
              <Dropdown.Menu className="mt-3 z-[1] w-52 menu-md">
                <Dropdown.Item onClick={logout}>Logout</Dropdown.Item>
              </Dropdown.Menu>
            </Dropdown>
             
          </Navbar.End>
        </Navbar>

        <div className={"p-8 " + props.className}>
          <div>{props.children}</div>
        </div>

      </Drawer>
    </div>
  );
};

const NavigationSidebar = (props: { className: string, logout: () => void }) => {
  return (
    <div className={`py-8 bg-base-100 w-72 h-screen overflow-y-scroll border-r border-gray-400 ${props.className}`}>
      <div className="mb-8 px-6 flex items-center justify-center">
        <Image src={LogoRect} alt="logo" className="max-h-[28px] " />
      </div>
      
      <Menu className="rounded-box">
        <Menu.Item>
        <Link href="/app" color="ghost"><HiOutlineChartBar className="w-[24px] h-[24px]" /> Dashboard</Link>
          <Link href="/app/loan" color="ghost"><HiOutlineCreditCard className="w-[24px] h-[24px]" /> Pinjaman (Loan)</Link>
        </Menu.Item>
        
      </Menu>
    </div>
    )
}

export default DashboardLayout;

