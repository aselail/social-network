import React from 'react'
import {Theme} from '@radix-ui/themes'
import LeftMenu from '@/components/LeftMenu'
import AddPost from '@/components/AddPost'
import Feed from '@/components/Feed'
import RightMenu from '@/components/RightMenu'

export default function HomePage() {
  return (
    <Theme appearance="light">
     <div className="flex gap-6">
      <div className="hidden xl:block" style={{ width: '20%' }}>
        <LeftMenu />
      </div>
      <div className="w-full lg:w-[70%] lg:w-[50%]">
        <div className="flex flex-col gap-6">
          <AddPost />
          <Feed />
        </div>
      </div>
      <div className="hidden xl:block" style={{ width: '30%' }}>
        <RightMenu />
      </div>
    </div>
    </Theme>
  )
}
