type AppHeaderProps = {
  title: string,
}

const AppHeader = ({title}: AppHeaderProps) => {
  return (
    <>
      <h1>{title}</h1>
      <hr/>
    </>
  )
}

export default AppHeader
