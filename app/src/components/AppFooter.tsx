import Grid from "@react-css/grid"


type AppFooterProps = {}

const AppFooter = ({}: AppFooterProps) => {
  return (
    <>
      <Grid className="footer">
        <p data-testid="footer">Copyright &copy; Uniquode</p>
      </Grid>
    </>
  )
}

export default AppFooter
