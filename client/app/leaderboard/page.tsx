import ProfileCard from "@/components/leaderboard/ProfileCard";
import { ProfileBadge } from "@/constants/enum";

// Mock data
const mockProfile = [
	{
		name: 'John Doe',
		avatar: 'https://picsum.photos/id/92/200/200',
		followers: 1223422,
		answeredQuestions: 1992,
		postedTutorials: 392,
		isFollowing: true,
		badges: [
			ProfileBadge.TOP_FOLLOWING,
			ProfileBadge.RISING_TALENT
		],
		latestPost: {
			title: "What's the best practice to learn Golang?",
			url: '/',
		},
		latestAnswer: {
			title: 'How to add swagger doc to the Gin app?',
			url: '/',
		}
	},
	{
		name: 'Alex Hunter',
		avatar: 'https://picsum.photos/id/93/200/200',
		followers: 1203917,
		isFollowing: false,
		badges: [
			ProfileBadge.TOP_FOLLOWING,
			ProfileBadge.TOP_LIKES,
			ProfileBadge.RISING_TALENT
		],
		answeredQuestions: 2392,
		postedTutorials: 523,
		latestPost: {
			title: "What's the best practice to learn Golang?",
			url: '/',
		},
		latestAnswer: {
			title: 'How to add swagger doc to the Gin app?',
			url: '/',
		}
	},
	{
		name: 'Armando Russo',
		avatar: 'https://picsum.photos/id/94/200/200',
		followers: 1923734,
		isFollowing: true,
		badges: [
			ProfileBadge.TOP_FOLLOWING,
			ProfileBadge.RISING_TALENT
		],
		answeredQuestions: 2271,
		postedTutorials: 102,
		latestPost: {
			title: "What's the best practice to learn Golang?",
			url: '/',
		},
		latestAnswer: {
			title: 'How to add swagger doc to the Gin app?',
			url: '/',
		}
	},
	{
		name: 'Scott Johnson',
		avatar: 'https://picsum.photos/id/95/200/200',
		followers: 1059234,
		isFollowing: false,
		badges: [
			ProfileBadge.TOP_FOLLOWING,
			ProfileBadge.TOP_LIKES,
			ProfileBadge.RISING_TALENT
		],
		answeredQuestions: 6721,
		postedTutorials: 392,
		latestPost: {
			title: "What's the best practice to learn Golang?",
			url: '/',
		},
		latestAnswer: {
			title: 'How to add swagger doc to the Gin app?',
			url: '/',
		}
	},
];

function LeaderBoardPage() {
	return (
		<div className="relative flex flex-col">
			<h1 className="bg-zinc-950 text-lg font-semibold">All Time Top Contributors</h1>
			<div className="mt-4 grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-8">
				{mockProfile.map((profile, index) => (
					<ProfileCard
						key={index}
						name={profile.name}
						avatar={profile.avatar}
						followers={profile.followers}
						isFollowing={profile.isFollowing}
						badges={profile.badges}
						answeredQuestions={profile.answeredQuestions}
						postedTutorials={profile.postedTutorials}
						latestAnswer={profile.latestAnswer}
						latestPost={profile.latestPost}
					/>
				))}
			</div>
			<h1 className="mt-16 text-lg font-semibold">Monthly Top Contributors</h1>
			<div className="mt-6 grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-8">
				{mockProfile.map((profile, index) => (
					<ProfileCard
						key={index}
						name={profile.name}
						avatar={profile.avatar}
						followers={profile.followers}
						isFollowing={profile.isFollowing}
						badges={profile.badges}
						answeredQuestions={profile.answeredQuestions}
						postedTutorials={profile.postedTutorials}
						latestAnswer={profile.latestAnswer}
						latestPost={profile.latestPost}
					/>
				))}
			</div>
			<h1 className="mt-16 text-lg font-semibold">Weekly Top Contributors</h1>
			<div className="mt-6 grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-8">
				{mockProfile.map((profile, index) => (
					<ProfileCard
						key={index}
						name={profile.name}
						avatar={profile.avatar}
						followers={profile.followers}
						isFollowing={profile.isFollowing}
						badges={profile.badges}
						answeredQuestions={profile.answeredQuestions}
						postedTutorials={profile.postedTutorials}
						latestAnswer={profile.latestAnswer}
						latestPost={profile.latestPost}
					/>
				))}
			</div>
		</div>
	);
}

export default LeaderBoardPage;