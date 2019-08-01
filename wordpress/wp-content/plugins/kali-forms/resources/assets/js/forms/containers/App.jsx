import supressNotices from '@/forms/utils/notices';
import AppBar from '@material-ui/core/AppBar';
import IconButton from '@material-ui/core/IconButton';
import InputAdornment from '@material-ui/core/InputAdornment';
import InputBase from '@material-ui/core/InputBase';
import Snackbar from '@material-ui/core/Snackbar';
// import Loader from './../components/Loader';
import { makeStyles } from '@material-ui/core/styles';
import Tab from '@material-ui/core/Tab';
import Tabs from '@material-ui/core/Tabs';
import Toolbar from '@material-ui/core/Toolbar';
import Tooltip from '@material-ui/core/Tooltip';
import Typography from '@material-ui/core/Typography';
import CloseIcon from '@material-ui/icons/Close';
// import EditIcon from '@material-ui/icons/Edit';
import SaveIcon from '@material-ui/icons/Save';
import React, { useEffect, useState } from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import * as PageActions from '../store/actions';
import ErrorBoundary from './ErrorBoundary';
import FormBuilder from './FormBuilder';
import FormSettings from './FormSettings';
import LogoSvg from './../../../img/icon.svg';
import TextField from '@material-ui/core/TextField';

const useStyles = makeStyles(theme => {
	const backgroundColorDefault =
		theme.palette.type === 'light' ? theme.palette.grey[100] : theme.palette.grey[900];

	return {
		app: {
			background: theme.palette.type === 'light' ? theme.palette.common.white : theme.palette.grey[800]
		},
		formNameInput: {
			marginRight: 50,
			flexGrow: 1,
			display: 'flex',
			'& > .MuiInputBase-root': {
				color: theme.palette.getContrastText(theme.palette.primary.main),
				background: theme.palette.primary.main,
				fontSize: 24,
				zIndex: 1,
				'&:before': {
					content: 'none'
				},
				'&:after': {
					content: 'none'
				}
			}
		},
		inputAdornment: {
			background: theme.palette.grey[900],
			height: '42px',
			borderRadius: '50%',
			width: '42px',
			alignItems: 'center',
			alignContent: 'center',
			justifyContent: 'center',
			marginRight: '15px',
		},
		logo: {
			maxWidth: '42px',
			marginRight: '15px',
		}
	}
})

const mapStateToProps = state => {
	return {
		loading: state.PageLoading,
		formInfo: state.FormInfo,
		ui: state.Ui,
		notices: supressNotices()
	};
};

const mapDispatchToProps = (dispatch) => {
	return bindActionCreators(PageActions, dispatch);
};


/**
 * App created as a hook
 *
 * @param {*} props
 * @returns
 */
const App = (props) => {
	const [activeTab, setActiveTab] = useState(props.ui.appBar);
	const [formName, setFormName] = useState(KaliFormsObject.formName);
	const [notices, setNotices] = useState(props.notices);
	const [loaded, setLoaded] = useState(false);
	/**
	 * Handle click
	 */
	const handleClick = (event, name) => {
		let action = name.toLowerCase();
		switch (action) {
			case 'delete':
				document.querySelector('#delete-action a').click();
				break;
			case 'save':
				document.getElementById('publish').click();
				break;
			case 'add-new':
				document.querySelector('.page-title-action').click();
				break;
			default:
				break;
		}
	};

	/**
	 * Tab toggler
	 * @param tab
	 */
	const toggle = (event, tab) => {
		if (activeTab !== tab) {
			setActiveTab(tab)
			props.setUiAppBar(tab)
			props.setTemplateSelectingFalse()
		}
	}

	/**
	 * Changes the form name
	 */
	const changeFormName = (event) => {
		let val = event.target.value;
		document.querySelector('#title').value = val
		setFormName(val);
	}

	const handleClose = (id) => {
		let newNotices = notices.filter((e, idx) => {
			return e.id !== id;
		});
		setNotices(newNotices);
	}

	useEffect(() => {
		if (props.formInfo.formName !== formName) {
			setFormName(props.formInfo.formName)
		}
	}, [props.formInfo.formName])

	useEffect(() => {
		setFormName(KaliFormsObject.formName)
	}, [])
	/**
	 * Display styles based on tab
	 */
	const displayStyles = {
		formBuilder: activeTab === 'formBuilder'
			? 'block' : 'none',
		formSettings: activeTab === 'formSettings'
			? 'block' : 'none',
	}

	const classes = useStyles()

	return (
		<div className={classes.app + ' kaliforms-wrapper'}>
			{/* <If condition={!loaded}>
				<Loader />
			</If> */}

			<ErrorBoundary>
				<AppBar position="static" color="primary" elevation={0} id="kali-appbar">
					<Toolbar>
						<img src={LogoSvg} className={classes.logo} />
						{/* <InputBase
							value={formName}
							onChange={e => changeFormName(e)}
							className={classes.formNameInput}
							placeholder={KaliFormsObject.translations.appBar.formName}
							startAdornment={
								<InputAdornment className={classes.inputAdornment}>
									<EditIcon />
								</InputAdornment>
							}
						/> */}
						<TextField
							value={formName}
							onChange={e => changeFormName(e)}
							className={classes.formNameInput}
							placeholder={KaliFormsObject.translations.appBar.formName}
						/>

						<Tabs indicatorColor="primary" value={activeTab} onChange={toggle}>
							<Tab value="formBuilder" label={KaliFormsObject.translations.appBar.formBuilder} />
							<Tab value="formSettings" label={KaliFormsObject.translations.appBar.formSettings} />
						</Tabs>
						<div style={{ marginLeft: 15 }}>
							<Tooltip title={KaliFormsObject.translations.general.save}>
								<IconButton
									onClick={(e) => handleClick(e, 'save')}
									color="inherit"
								>
									<SaveIcon />
								</IconButton>
							</Tooltip>
							<Tooltip title={KaliFormsObject.translations.appBar.backToWp}>
								<IconButton
									onClick={() => window.location.href = KaliFormsObject.exit_url}
									color="inherit"
								>
									<CloseIcon />
								</IconButton>
							</Tooltip>
						</div>
					</Toolbar>
				</AppBar>
				{
					notices.map(e => <Snackbar anchorOrigin={{ vertical: 'bottom', horizontal: 'left' }}
						open={true}
						message={e.message}
						key={e.id}
						action={[
							<IconButton
								key="close"
								aria-label="Close"
								color="inherit"
								onClick={() => handleClose(e.id)}
							>
								<CloseIcon />
							</IconButton>,
						]}
					/>)
				}
				<Typography component="div" style={{ display: displayStyles.formBuilder }}>
					<FormBuilder />
				</Typography>
				<Typography component="div" style={{ display: displayStyles.formSettings }}>
					<FormSettings />
				</Typography>
			</ErrorBoundary>
		</div>
	)
}

export default connect(
	mapStateToProps,
	mapDispatchToProps
)(App);

