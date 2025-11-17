import Image from "next/image";
import Comments from "./Comments";

const Post = () => {
  return (
    <div className='flex flex-col gap-4'>
      {/* USER */}
      <div className="flex items-center justify-between">
        <div className='flex items-center gap-2'>
          <img 
          src="https://images.pexels.com/photos/31740983/pexels-photo-31740983.jpeg" 
          width={40} height={40} 
          alt=""
          className="w-10 h-10 rounded-full"
          />
          <span className="font-medium">Ali Selail</span>
          </div>
          <Image src="/more.png" width={16} height={16} alt="" />
      </div>
      {/* DESC */}
      <div className="flex flex-col gap-4">
        <div className="w-full min-h-96 relative">
        <Image 
          src="https://images.pexels.com/photos/31740983/pexels-photo-31740983.jpeg" 
          fill
          alt="" 
          className="object-cover rounded-md"
        />
        </div>
        <p>Post Contant</p>
      </div>
      {/* INTRACTION */}
      <div className="flex item-center justify-between teaxt -sm my-4">
        <div className="flex gap-8">
          <div className='flex item-center gap-4 bg-slate-50 p-2 rounded-xl'>
            <Image src="/like.png" width={16} height={16} alt="" className="crusor-pointer" />
            <span className="text-gray-300"></span>
            <span className="text-gray-500">12
              <span className="hidden md:inline"> Likes</span>
            </span>
          </div>
           <div className='flex item-center gap-4 bg-slate-50 p-2 rounded-xl'>
            <Image src="/comment.png" width={16} height={16} alt="" className="crusor-pointer" />
            <span className="text-gray-300"></span>
            <span className="text-gray-500">12
              <span className="hidden md:inline"> Comments</span>
            </span>
          </div>
        </div>
        <div className="">
          <div className='flex item-center gap-4 bg-slate-50 p-2 rounded-xl'>
            <Image src="/Share.png" width={16} height={16} alt="" className="crusor-pointer" />
            <span className="text-gray-300"></span>
            <span className="text-gray-500">12
              <span className="hidden md:inline"> Share</span>
            </span>
          </div>
        </div>
      </div>
      <Comments/>
    </div>
  );
};

export default Post;