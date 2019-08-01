import React from 'react';
import Card from '@material-ui/core/Card';
import CardHeader from '@material-ui/core/CardHeader';
import IconButton from '@material-ui/core/IconButton';
import AddIcon from '@material-ui/icons/Add';
import { bindActionCreators } from 'redux';
import connect from 'react-redux/es/connect/connect';
import * as StoreActions from './../../store/actions'

const mapStateToProps = state => {
	return {
		loading: state.PageLoading,
		formFieldEditor: state.FormFieldEditor,
		fieldComponents: state.FieldComponents,
		sidebar: state.Sidebar,
		fieldComponentProperties: formatFieldCollection(state.Sidebar.fieldComponents)
	};
};

const formatFieldCollection = collection => {
	const fields = {};
	collection.map(group => {
		group.fields.map(field => {
			fields[field.id] = field.properties;
		})
	})
	return fields;
}

const mapDispatchToProps = (dispatch) => {
	return bindActionCreators(StoreActions, dispatch);
};

const styles = {
	display: 'flex',
	alignItems: 'center',
	width: '100%',
	height: '100%',
	paddingLeft: 15,
}

const SidebarFieldComponentItem = (props) => {
	/**
	 * Add a field in the builder
	 * SHOULD BE PARSED/STRINGIFY SO YOU FORGET OBJECT REFS
	 */
	const addField = () => {
		props.properties.id.value = props.properties.id.value + props.fieldComponents.length
		props.addComponent(
			JSON.parse(JSON.stringify(
				{
					id: props.id,
					label: props.label,
					properties: formatObj(props),
					constraint: props.constraint,
					internalId: props.id.toLowerCase() + props.fieldComponents.length
				}
			))
		);
	}
	/**
	 * Format object
	 *
	 * @param {*} obj
	 * @returns
	 * @memberof SidebarFieldComponentItem
	 */
	const formatObj = (obj) => {
		let properties = {};
		for (const key in obj.properties) {
			properties[key] = obj.properties[key].value
		}

		return properties;
	}

	return (
		<Card>
			<CardHeader
				style={{ padding: '6px 16px' }}
				titleTypographyProps={{ variant: 'subtitle2' }}
				action={
					<span style={{ position: 'relative', top: 3 }}>
						<IconButton onClick={() => addField()}>
							<AddIcon />
						</IconButton>
					</span>
				}
				title={props.label}
			/>
		</Card>
	)
}

export default connect(
	mapStateToProps,
	mapDispatchToProps
)(SidebarFieldComponentItem);
