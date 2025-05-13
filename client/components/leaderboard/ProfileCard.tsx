import { ProfileBadge } from '@/constants/enum';
import { formatNumberWithUnits } from '@/lib/utils';
import React from 'react';
import { FaBook, FaChartLine, FaComments, FaRegStar, FaStar } from 'react-icons/fa';

interface ProfileCardProps {
  name: string;
  avatar: string;
  followers: number;
  isFollowing?: boolean;
  badges: ProfileBadge[];
  answeredQuestions: number;
  postedTutorials: number;
  latestAnswer: {
    title: string;
    url: string;
  };
  latestPost: {
    title: string;
    url: string;
  }
}

const ProfileCard: React.FC<ProfileCardProps> = ({
  name,
  avatar,
  followers,
  isFollowing,
  badges,
  answeredQuestions,
  postedTutorials,
  latestAnswer,
  latestPost
}) => {
  return (
    <div className="relative flex flex-col items-center gap-6 rounded-lg border border-zinc-700 border-dashed py-8 px-6">
      {isFollowing ?
        <FaStar size={18} className="absolute top-6 right-6 text-xs text-amber-300" />
        : <FaRegStar size={18} className="absolute top-6 right-6 text-xs text-zinc-600" />
      }
      <div className="flex flex-col items-center">
        <img
          src={avatar}
          alt={`${name}'s avatar`}
          className="w-28 h-28 rounded-full"
        />
        <h2 className="mt-2 text-lg font-semibold">{name}</h2>
        <p className="text-xs text-gray-500">{formatNumberWithUnits(followers)} Followers</p>
      </div>
      <div className="flex flex-row gap-4">
        <p className="flex flex-row items-center gap-2 text-sm">
          <FaComments />{answeredQuestions}
        </p>
        <p className="flex flex-row items-center gap-2 text-sm">
          <FaBook />{postedTutorials}
        </p>
      </div>
      {/* Badges */}
      <div className='text-[24px] flex flex-row gap-2'>
        {badges.findIndex((badge) => badge === ProfileBadge.TOP_FOLLOWING) >= 0 &&
          <p className="cursor-pointer border border-zinc-700 p-2 rounded-lg" title="Top Following">🤩</p>
        }
        {badges.findIndex((badge) => badge === ProfileBadge.TOP_LIKES) >= 0 &&
          <p className="cursor-pointer border border-zinc-700 p-2 rounded-lg" title="Top Likes">🌟</p>
        }
        {badges.findIndex((badge) => badge === ProfileBadge.RISING_TALENT) >= 0 &&
          <p className="cursor-pointer border border-zinc-700 p-2 rounded-lg" title="Rising Talents">🚀</p>
        }
      </div>
      <div className="flex flex-col gap-1 w-full">
        <p className="text-[13px] text-zinc-400">Latest Post:</p>
        <a href={latestPost.url} className="text-[13px] hover:underline">{latestPost.title}</a>
      </div>
      <div className="flex flex-col gap-1 w-full">
        <p className="text-[13px] text-zinc-400">Latest Answer:</p>
        <a href={latestAnswer.url} className="text-[13px] hover:underline">{latestAnswer.title}</a>
      </div>
    </div>
  );
};

export default ProfileCard;