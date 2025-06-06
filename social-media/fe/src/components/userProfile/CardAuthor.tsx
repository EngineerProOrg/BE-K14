import { Avatar, CardHeader } from "@mui/material";

interface CardAuthorProps {
  avatarSrc: string;
  title: string;
  subheader: string;
}

function CardAuthor({ avatarSrc, title, subheader }: CardAuthorProps) {
  return (
    <>
      <CardHeader
        avatar={<Avatar alt={title} src={avatarSrc} />}
        title={title}
        subheader={subheader}
      />
    </>
  );
}

export default CardAuthor;
