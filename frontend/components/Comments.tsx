import { DividerHorizontalIcon } from "@radix-ui/react-icons";
import Image from "next/image";

const Comments = () => {
    return (
        <div className="">
            {/* WRITE */}
            <div className="flex items-center gap-4">
                <img src="https://images.pexels.com/photos/31740983/pexels-photo-31740983.jpeg" 
                width={32} height={32} 
                alt=""
                className=" w-8 h-8 rounded-full flex-1"
                /> 
                <div className="flex-1 flex items-center justify-between bg-slate-100 text-sm px-6 w-full">
                    <input type="text" placeholder="Write a comment..." className="bg-transparent outline-none flex-1" />
                    <Image src="/emoji.png" width={16} height={16} alt="" className="cursor-pointer" />
                </div>
            </div>
            {/* COMMENTS */}
            <div className=''>
                {/* COMMENT */}
                <div className="">
                    {/* AVATAR */ }
                    <Image src="https://images.pexels.com/photos/7651438/pexels-photo-7651438.jpeg"
                     alt="" 
                    width={40} 
                    height={40} 
                    className="w-10 h-10 rounded-full"
                     />
                     {/* DESC */}
                     <div className="flex flex-col gap-2">
                        <span className="font-medium">Yousif Ahmed</span>
                        <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Animi sint harum molestiae sit veniam exercitationem consectetur, dolorem culpa accusantium, sapiente facilis error itaque aut commodi iste adipisci illum. Quia, expedita!</p>
                        <div className="">
                            <div className="">
                                <Image src="/like.png"
                                 alt=""
                                  width={12}
                                    height={12}
                                     className="crusor-pointer w-4 h-4"
                                      />
                                    <span className="text-gray-300">|</span>
                                    <span className="text-gray-500">12 Likes</span>
                            </div>
                        </div>
                     </div>
                     {/* ICON */}
                     <Image src="/more.png"
                     alt=""
                     width={16}
                     height={16}
                     className="cursor-pointer w-4 h-4"
                     ></Image>
                </div>
            </div>
        </div>
    );
};

export default Comments;